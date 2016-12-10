package maths

import (
	"github.com/gonum/matrix/mat64"
	"github.com/gonum/stat"
)

func (vX *Var) Regress(vY *Var) [2]float64 {
	c, b := stat.LinearRegression(vX.Data, vY.Data, nil, false)
	return [2]float64{c, b}
}
func (vX *Var) Predictions(line [2]float64) (vP *Var) {
	vP = NewVar("Predictions")
	vP.Data = make([]float64, len(vX.Data))
	for i, x := range vX.Data {
		vP.Data[i] = line[0] + line[1]*x
	}
	vP.CalcMeanStd()
	return vP
}
func (vX *Var) Residuals(vY *Var, line [2]float64) (vR *Var) {
	vR = NewVar("Residual " + vY.Name)
	vR.Data = make([]float64, len(vX.Data))
	for i, x := range vX.Data {
		vR.Data[i] = vY.Data[i] - (line[0] + line[1]*x)
	}
	vR.CalcMeanStd()
	return vR
}

// Regress returns n+1 len float64 slice, with first
// value being the intercept term of the regression model
func (vs Vars) Regress(vD *Var) ([]float64, error) {
	xM := vs.RegressionMatrix()
	yM := CollectVars(vD).Matrix()

	var qr mat64.QR
	qr.Factorize(xM)

	var bM mat64.Dense
	err := bM.SolveQR(&qr, false, yM)
	if err != nil {
		return nil, err
	}
	sol := make([]float64, len(vs)+1)
	mat64.Col(sol, 0, &bM)
	return sol, nil
}
