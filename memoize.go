package memoize

import (
	"github.com/AlexandreChamard/go-cache"
)

type memoizerCacheInterface[K comparable, V any] interface {
	Set(K, V)
	Get(K) (V, bool)
}

type memoizerDefaultCache[K comparable, V any] map[K]V

func (this memoizerDefaultCache[K, V]) Set(k K, v V) {
	this[k] = v
}

func (this memoizerDefaultCache[K, V]) Get(k K) (V, bool) {
	v, ok := this[k]
	return v, ok
}

// Returns: the memoized function and its cancel function when the function is no longer used (cleanup)
func Memoize1[Ret any, T1 comparable](f func(T1) Ret, options ...*memoizationOption) func(T1) Ret {
	var memoizer memoizerCacheInterface[T1, Ret]

	if len(options) > 0 {
		option := options[0]
		memoizer = cache.NewCache[T1, Ret](cache.NewCacheOptions().
			MaxSize(int(option.maxMemoize)).
			CacheDuration(option.expiration))
	} else {
		memoizer = memoizerDefaultCache[T1, Ret]{}
	}

	return func(t T1) Ret {
		if ret, ok := memoizer.Get(t); ok {
			return ret
		}
		ret := f(t)
		memoizer.Set(t, ret)
		return ret
	}
}

func Memoize2[Ret any, T1, T2 comparable](f func(T1, T2) Ret, options ...*memoizationOption) func(T1, T2) Ret {
	var memoizer memoizerCacheInterface[struct {
		t1 T1
		t2 T2
	}, Ret]

	if len(options) > 0 {
		option := options[0]
		memoizer = cache.NewCache[struct {
			t1 T1
			t2 T2
		}, Ret](cache.NewCacheOptions().
			MaxSize(int(option.maxMemoize)).
			CacheDuration(option.expiration))
	} else {
		memoizer = memoizerDefaultCache[struct {
			t1 T1
			t2 T2
		}, Ret]{}
	}

	return func(t1 T1, t2 T2) Ret {
		args := struct {
			t1 T1
			t2 T2
		}{t1, t2}

		if ret, ok := memoizer.Get(args); ok {
			return ret
		}
		ret := f(t1, t2)
		memoizer.Set(args, ret)
		return ret
	}
}

func Memoize3[Ret any, T1, T2, T3 comparable](f func(T1, T2, T3) Ret, options ...*memoizationOption) func(T1, T2, T3) Ret {
	var memoizer memoizerCacheInterface[struct {
		t1 T1
		t2 T2
		t3 T3
	}, Ret]

	if len(options) > 0 {
		option := options[0]
		memoizer = cache.NewCache[struct {
			t1 T1
			t2 T2
			t3 T3
		}, Ret](cache.NewCacheOptions().
			MaxSize(int(option.maxMemoize)).
			CacheDuration(option.expiration))
	} else {
		memoizer = memoizerDefaultCache[struct {
			t1 T1
			t2 T2
			t3 T3
		}, Ret]{}
	}

	return func(t1 T1, t2 T2, t3 T3) Ret {
		args := struct {
			t1 T1
			t2 T2
			t3 T3
		}{t1, t2, t3}

		if ret, ok := memoizer.Get(args); ok {
			return ret
		}
		ret := f(t1, t2, t3)
		memoizer.Set(args, ret)
		return ret
	}
}

func Memoize4[Ret any, T1, T2, T3, T4 comparable](f func(T1, T2, T3, T4) Ret, options ...*memoizationOption) func(T1, T2, T3, T4) Ret {
	var memoizer memoizerCacheInterface[struct {
		t1 T1
		t2 T2
		t3 T3
		t4 T4
	}, Ret]

	if len(options) > 0 {
		option := options[0]
		memoizer = cache.NewCache[struct {
			t1 T1
			t2 T2
			t3 T3
			t4 T4
		}, Ret](cache.NewCacheOptions().
			MaxSize(int(option.maxMemoize)).
			CacheDuration(option.expiration))
	} else {
		memoizer = memoizerDefaultCache[struct {
			t1 T1
			t2 T2
			t3 T3
			t4 T4
		}, Ret]{}
	}

	return func(t1 T1, t2 T2, t3 T3, t4 T4) Ret {
		args := struct {
			t1 T1
			t2 T2
			t3 T3
			t4 T4
		}{t1, t2, t3, t4}

		if ret, ok := memoizer.Get(args); ok {
			return ret
		}
		ret := f(t1, t2, t3, t4)
		memoizer.Set(args, ret)
		return ret
	}
}

func Memoize5[Ret any, T1, T2, T3, T4, T5 comparable](f func(T1, T2, T3, T4, T5) Ret, options ...*memoizationOption) func(T1, T2, T3, T4, T5) Ret {
	var memoizer memoizerCacheInterface[struct {
		t1 T1
		t2 T2
		t3 T3
		t4 T4
		t5 T5
	}, Ret]

	if len(options) > 0 {
		option := options[0]
		memoizer = cache.NewCache[struct {
			t1 T1
			t2 T2
			t3 T3
			t4 T4
			t5 T5
		}, Ret](cache.NewCacheOptions().
			MaxSize(int(option.maxMemoize)).
			CacheDuration(option.expiration))
	} else {
		memoizer = memoizerDefaultCache[struct {
			t1 T1
			t2 T2
			t3 T3
			t4 T4
			t5 T5
		}, Ret]{}
	}

	return func(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5) Ret {
		args := struct {
			t1 T1
			t2 T2
			t3 T3
			t4 T4
			t5 T5
		}{t1, t2, t3, t4, t5}

		if ret, ok := memoizer.Get(args); ok {
			return ret
		}
		ret := f(t1, t2, t3, t4, t5)
		memoizer.Set(args, ret)
		return ret
	}
}
