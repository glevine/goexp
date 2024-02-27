package main

import (
	"fmt"

	mathv2 "github.com/glevine/goexp/interfaces/v2/math"
	addv2 "github.com/glevine/goexp/interfaces/v2/math/add"
	subv2 "github.com/glevine/goexp/interfaces/v2/math/subtract"
	mathv3 "github.com/glevine/goexp/interfaces/v3/math"
	addv3 "github.com/glevine/goexp/interfaces/v3/math/add"
	subv3 "github.com/glevine/goexp/interfaces/v3/math/subtract"
)

func main() {
	/*
	 * Note: Functions are not a substitute for interfaces, even if the signatures are identical.
	 *
	 * import "github.com/glevine/goexp/interfaces/math"
	 */
	// adderv1 := math.NewOperator(add.NewOperator())
	// fmt.Println(adderv1(2, 2))

	// subberv1 := math.NewOperator(subtract.NewOperator())
	// fmt.Println(subberv1(2, 2))

	/*
	 * Note: Referencing the same function type from different implementations works ok.
	 */
	adderv2 := mathv2.NewOperator(addv2.NewOperator())
	fmt.Println(adderv2(2, 2))

	subberv2 := mathv2.NewOperator(subv2.NewOperator())
	fmt.Println(subberv2(5, 1))

	/*
	 * Note: Polymorphism is really only achieved through true interfaces. But this is only needed
	 * if a single type needs to be capable of satisfying multiple interfaces.
	 */
	adderv3 := mathv3.NewService(addv3.NewService())
	fmt.Println(adderv3.Operate(6, 3))

	subberv3 := mathv3.NewService(subv3.NewService())
	fmt.Println(subberv3.Operate(7, 4))
}
