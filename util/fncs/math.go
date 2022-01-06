// Copyright © 2016 SurrealDB Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fncs

import (
	"context"

	"github.com/surrealdb/surrealdb/util/math"
)

func mathAbs(ctx context.Context, args ...interface{}) (interface{}, error) {
	if val, ok := ensureFloat(args[0]); ok {
		return outputFloat(math.Abs(val))
	}
	return nil, nil
}

func mathBottom(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	if take, ok := ensureInt(args[1]); ok {
		return math.Bottom(vals, int(take)), nil
	}
	return nil, nil
}

func mathCeil(ctx context.Context, args ...interface{}) (interface{}, error) {
	if val, ok := ensureFloat(args[0]); ok {
		return outputFloat(math.Ceil(val))
	}
	return nil, nil
}

func mathCorrelation(ctx context.Context, args ...interface{}) (interface{}, error) {
	a := ensureFloats(args[0])
	b := ensureFloats(args[1])
	return outputFloat(math.Correlation(a, b))
}

func mathCovariance(ctx context.Context, args ...interface{}) (interface{}, error) {
	a := ensureFloats(args[0])
	b := ensureFloats(args[1])
	return outputFloat(math.Covariance(a, b))
}

func mathFixed(ctx context.Context, args ...interface{}) (interface{}, error) {
	if val, ok := ensureFloat(args[0]); ok {
		if pre, ok := ensureInt(args[1]); ok {
			return outputFixed(val, pre)
		}
	}
	return nil, nil
}

func mathFloor(ctx context.Context, args ...interface{}) (interface{}, error) {
	if val, ok := ensureFloat(args[0]); ok {
		return outputFloat(math.Floor(val))
	}
	return nil, nil
}

func mathGeometricmean(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	return outputFloat(math.GeometricMean(vals))
}

func mathHarmonicmean(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	return outputFloat(math.HarmonicMean(vals))
}

func mathInterquartile(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	return outputFloat(math.InterQuartileRange(vals))
}

func mathMax(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	return outputFloat(math.Max(vals))
}

func mathMean(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	return outputFloat(math.Mean(vals))
}

func mathMedian(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	return outputFloat(math.Median(vals))
}

func mathMidhinge(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	return outputFloat(math.Midhinge(vals))
}

func mathMin(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	return outputFloat(math.Min(vals))
}

func mathMode(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	return math.Mode(vals), nil
}

func mathNearestRank(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	if perc, ok := ensureFloat(args[1]); ok {
		return outputFloat(math.NearestRankPercentile(vals, perc))
	}
	return nil, nil
}

func mathPercentile(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	if perc, ok := ensureFloat(args[1]); ok {
		return outputFloat(math.Percentile(vals, perc))
	}
	return nil, nil
}

func mathRound(ctx context.Context, args ...interface{}) (interface{}, error) {
	if val, ok := ensureFloat(args[0]); ok {
		return outputFloat(math.Round(val))
	}
	return nil, nil
}

func mathSample(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	if take, ok := ensureInt(args[1]); ok {
		return math.Sample(vals, int(take)), nil
	}
	return nil, nil
}

func mathSpread(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	return outputFloat(math.Spread(vals))
}

func mathSqrt(ctx context.Context, args ...interface{}) (interface{}, error) {
	if val, ok := ensureFloat(args[0]); ok {
		return outputFloat(math.Sqrt(val))
	}
	return nil, nil
}

func mathStddev(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	return outputFloat(math.SampleStandardDeviation(vals))
}

func mathSum(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	return outputFloat(math.Sum(vals))
}

func mathTop(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	if take, ok := ensureInt(args[1]); ok {
		return math.Top(vals, int(take)), nil
	}
	return nil, nil
}

func mathTrimean(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	return outputFloat(math.Trimean(vals))
}

func mathVariance(ctx context.Context, args ...interface{}) (interface{}, error) {
	vals := ensureFloats(args[0])
	return outputFloat(math.SampleVariance(vals))
}