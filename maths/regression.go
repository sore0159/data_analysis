package maths

type Regression struct {
	IVars Vars
	DVar  *Var
	Coef  []float64
	//
	N           float64
	R           float64
	Predictions *Var
	Residuals   *Var
	ResSqSum    float64
}

func FullRegression(vD *Var, vs Vars) (*Regression, error) {
	coef, err := vs.Regress(vD)
	if err != nil {
		return nil, err
	}
	r := &Regression{
		IVars:       vs,
		DVar:        vD,
		Coef:        coef,
		N:           float64(len(vD.Data)),
		R:           float64(len(coef)),
		Predictions: vs.Predictions(coef),
	}
	r.Residuals = vD.Copy()
	r.Residuals.Add(r.Predictions, -1)
	r.Residuals.Name = "Residuals From Prediction"
	r.ResSqSum = r.Residuals.SumSqs()

	return r, nil
}

func (r Regression) MeanSqError() float64 {
	return r.ResSqSum / (r.N - r.R)
}
func (r Regression) MeanSqResid() float64 {
	return r.ResSqSum / r.N
}
func (r Regression) RSq() float64 {
	return 1 - (r.ResSqSum / r.DVar.SumSqs())
}

func (r Regression) LineFor(x int) [2]float64 {
	c := r.Coef[0]
	for j, v2 := range r.IVars {
		mn, _ := v2.CalcMeanStd() // should be just zero
		if x != j {
			c += r.Coef[j+1] * mn
		}
	}
	return [2]float64{c, r.Coef[x+1]}
}
