// Copyright (C) 2024 Tim Bastin, l3montree UG (haftungsbeschränkt)
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package utils

import (
	"sync"

	"golang.org/x/sync/errgroup"
)

type concurrentResult struct {
	index int
	value any
}

type errGroup[T any] struct {
	ch             chan T
	group          errgroup.Group
	res            []T
	collectionDone sync.WaitGroup
}

func (eg *errGroup[T]) Go(fn func() (T, error)) {
	eg.group.Go(func() error {
		r, err := fn()
		if err != nil {
			return err
		}

		eg.ch <- r

		return nil
	})
}

func (eg *errGroup[T]) startCollecting() {
	// reset the result slice
	eg.res = make([]T, 0)
	eg.collectionDone.Add(1)
	go func() {
		defer eg.collectionDone.Done()
		for r := range eg.ch {
			eg.res = append(eg.res, r)
		}
	}()
}

func (eg *errGroup[T]) SetLimit(limit int) {
	eg.group.SetLimit(limit)
}

func (eg *errGroup[T]) WaitAndCollect() ([]T, error) {
	defer eg.startCollecting()
	err := eg.group.Wait()
	close(eg.ch)
	// Wait for the collection to finish - otherwise the result might be incomplete
	eg.collectionDone.Wait()
	if err != nil {
		return nil, err
	}
	// Reset the channel
	eg.ch = make(chan T)
	res := eg.res

	return res, nil
}

func ErrGroup[T any](limit int) *errGroup[T] {
	g := errGroup[T]{ch: make(chan T), group: errgroup.Group{}}
	g.group.SetLimit(limit)
	g.startCollecting() // otherwise a call to .Go will block, since we are not reading from the channel
	return &g
}

func Concurrently(fns ...func() any) []any {
	results := make([]concurrentResult, len(fns))
	ch := make(chan concurrentResult, len(fns))
	for i, fn := range fns {
		go func(i int, fn func() any) {
			ch <- concurrentResult{i, fn()}
		}(i, fn)
	}
	for i := 0; i < len(fns); i++ {
		results[i] = <-ch
	}

	res := make([]any, len(fns))
	for _, r := range results {
		res[r.index] = r.value
	}

	return res
}
