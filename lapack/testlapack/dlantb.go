// Copyright ©2020 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testlapack

import (
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/lapack"
)

type Dlantber interface {
	Dlantb(norm lapack.MatrixNorm, uplo blas.Uplo, diag blas.Diag, n, k int, a []float64, lda int, work []float64) float64
}
