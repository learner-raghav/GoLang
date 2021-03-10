package entity

type Variant struct {
	Variant_id int `json:"variant_id"`
	Model_id int `json:"model_id"`
	Variant_name string `json:"variant_name"`
	Disp float64 `json:"disp"`
	Peak_power float64 `json:"peak_power"`
	Peak_torque float64 `json:"peak_torque"`
	Model_name string `json:"model_name"`
	Brand_name string `json:"brand_name"`
	Brand_Id int `json:"brand_id"`
}
