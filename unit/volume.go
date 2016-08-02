package unit

// ...
const (
	// SI
	CubicYoctometer        = CubicMeter * 1e-72
	CubicZeptometer        = CubicMeter * 1e-63
	CubicAttometer         = CubicMeter * 1e-54
	CubicFemtometer        = CubicMeter * 1e-45
	CubicPicometer         = CubicMeter * 1e-36
	CubicNanometer         = CubicMeter * 1e-27
	CubicMicrometer        = CubicMeter * 1e-18
	CubicMillimeter        = CubicMeter * 1e-9
	CubicCentimeter        = CubicMeter * 1e-6
	CubicDecimeter         = CubicMeter * 1e-3
	CubicMeter      Volume = 1e0
	CubicDecameter         = CubicMeter * 1e3
	CubicHectometer        = CubicMeter * 1e6
	CubicKilometer         = CubicMeter * 1e9
	CubicMegameter         = CubicMeter * 1e18
	CubicGigameter         = CubicMeter * 1e27
	CubicTerameter         = CubicMeter * 1e36
	CubicPetameter         = CubicMeter * 1e45
	CubicExameter          = CubicMeter * 1e54
	CubicZettameter        = CubicMeter * 1e63
	CubicYottameter        = CubicMeter * 1e72

	// SI derived
	Yoctoliter = Liter * 1e-24
	Zepoliter  = Liter * 1e-21
	Attoliter  = Liter * 1e-18
	Femtoliter = Liter * 1e-15
	Picoliter  = Liter * 1e-12
	Nanoliter  = Liter * 1e-9
	Microliter = Liter * 1e-6
	Milliliter = Liter * 1e-3
	Centiliter = Liter * 1e-2
	Deciliter  = Liter * 1e-1
	Liter      = CubicMeter * 1e-3
	Decaliter  = Liter * 1e1
	Hectoliter = Liter * 1e2
	Kiloliter  = Liter * 1e3
	Megaliter  = Liter * 1e6
	Gigaliter  = Liter * 1e9
	Teraliter  = Liter * 1e12
	Petaliter  = Liter * 1e15
	Exaliter   = Liter * 1e18
	Zettaliter = Liter * 1e21
	Yottaliter = Liter * 1e24

	// US
	CubicInch    = Liter * 0.016387064
	CubicFoot    = CubicInch * 1728
	CubicYard    = CubicFoot * 27
	CubicMile    = CubicYard * 5451776000
	CubicFurlong = CubicMile * 0.00195314

	// imperial liquid
	ImperialGallon     = Liter * 4.54609
	ImperialQuart      = ImperialGallon / 4
	ImperialPint       = ImperialQuart / 2
	ImperialCup        = ImperialPint / 2
	ImperialGill       = ImperialPint / 4
	ImperialFluidOunce = ImperialGill / 5
	ImperialFluidDram  = ImperialFluidOunce / 8
	ImperialPeck       = ImperialGallon * 2
	ImperialBushel     = ImperialPeck * 4

	// US liquid
	USLiquidGallon = CubicInch * 231
	USLiquidQuart  = CubicInch * 57.75
	USLiquidPint   = CubicInch * 28.875
	USCup          = USLiquidPint / 2
	USLegalCup     = Milliliter * 240
	USGill         = Milliliter * 118.29411825
	USTableSpoon   = USFluidOunce / 2
	USTeaSpoon     = USTableSpoon / 3
	USFluidDram    = USFluidOunce / 8
	USFluidOunce   = USLiquidGallon / 128

	// US dry
	USDryQuart  = USDryGallon / 4
	USBushel    = USPeck * 4
	USPeck      = USDryGallon * 2
	USDryGallon = CubicInch * 268.8025
	USDryPint   = CubicInch * 33.6003125
)
