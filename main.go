/*
 * GNU GLPK linear/mixed integer solver
 * short.c example on go
 */
package main

/*
#cgo CFLAGS: -I/usr/local/include/
#cgo LDFLAGS: -L/usr/local/lib -lglpk
#include <glpk.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import (
	"fmt"
)

const ()

var ()

func main() {
	fmt.Println("hello")

	/* declare variables */
	//glp_prob *lp;
	var lp *C.glp_prob
	//int ia[1+1000], ja[1+1000];
	var ia [1 + 1000]int
	var ja [1 + 1000]int
	//double ar[1+1000], z, x1, x2;
	var ar [1 + 1000]float64
	var z float64
	var x1 float64
	var x2 float64

	/* create problem */
	//lp = glp_create_prob();
	lp = C.glp_create_prob()
	// A common idiom in cgo programs is to defer the free immediately after allocating
	/* housekeeping */
	defer C.glp_delete_prob(lp)
	defer C.glp_free_env()

	//glp_set_prob_name(lp, "short");
	lp_probe_name := C.CString("short")
	defer C.free(unsafe.Pointer(lp_probe_name))
	_, err := C.glp_set_prob_name(lp, lp_probe_name)
	//glp_set_obj_dir(lp, GLP_MAX);
	_, err = C.glp_set_obj_dir(lp, C.GLP_MAX)

	/* fill problem */
	//glp_add_rows(lp, 2);
	_, err = C.glp_add_rows(lp, 2)
	//glp_set_row_name(lp, 1, "p");
	row_name1 := C.CString("p")
	defer C.free(unsafe.Pointer(row_name1))
	_, err = C.glp_set_row_name(lp, 1, row_name1)
	//glp_set_row_bnds(lp, 1, GLP_UP, 0.0, 1.0);
	_, err = C.glp_set_row_bnds(lp, 1, C.GLP_UP, 0.0, 1.0)
	//glp_set_row_name(lp, 2, "q");
	row_name2 := C.CString("q")
	defer C.free(unsafe.Pointer(row_name2))
	_, err = C.glp_set_row_name(lp, 2, row_name2)
	//glp_set_row_bnds(lp, 2, GLP_UP, 0.0, 2.0);
	_, err = C.glp_set_row_bnds(lp, 2, C.GLP_UP, 0.0, 2.0)
	//glp_add_cols(lp, 2);
	_, err = C.glp_add_cols(lp, 2)
	//glp_set_col_name(lp, 1, "x1");
	col_name1 := C.CString("x1")
	defer C.free(unsafe.Pointer(col_name1))
	_, err = C.glp_set_col_name(lp, 1, col_name1)
	//glp_set_col_bnds(lp, 1, GLP_LO, 0.0, 0.0);
	_, err = C.glp_set_col_bnds(lp, 1, C.GLP_LO, 0.0, 0.0)
	//glp_set_obj_coef(lp, 1, 0.6);
	_, err = C.glp_set_obj_coef(lp, 1, 0.6)
	//glp_set_col_name(lp, 2, "x2");
	col_name2 := C.CString("x2")
	defer C.free(unsafe.Pointer(col_name2))
	_, err = C.glp_set_col_name(lp, 2, col_name2)
	//glp_set_col_bnds(lp, 2, GLP_LO, 0.0, 0.0);
	_, err = C.glp_set_col_bnds(lp, 2, C.GLP_LO, 0.0, 0.0)
	//glp_set_obj_coef(lp, 2, 0.5);
	_, err = C.glp_set_obj_coef(lp, 2, 0.5)
	//ia[1] = 1, ja[1] = 1, ar[1] = 1.0; /* a[1,1] = 1 */
	//ia[2] = 1, ja[2] = 2, ar[2] = 2.0; /* a[1,2] = 2 */
	//ia[3] = 2, ja[3] = 1, ar[3] = 3.0; /* a[2,1] = 3 */
	//ia[4] = 2, ja[4] = 2, ar[4] = 1.0; /* a[2,2] = 1 */
	//glp_load_matrix(lp, 4, ia, ja, ar);
	ia[1] = 1
	ja[1] = 1
	ar[1] = 1.0 /* a[1,1] = 1 */
	ia[2] = 1
	ja[2] = 2
	ar[2] = 2.0 /* a[1,2] = 2 */
	ia[3] = 2
	ja[3] = 1
	ar[3] = 3.0 /* a[2,1] = 3 */
	ia[4] = 2
	ja[4] = 2
	ar[4] = 1.0 /* a[2,2] = 1 */
	_, err = C.glp_load_matrix(lp, 4, ia, ja, ar)

	/* solve problem */
	//glp_simplex(lp, NULL);
	_, err = C.glp_simplex(lp, nil)

	/* recover and display results */
	//z = glp_get_obj_val(lp);
	z, err := C.glp_get_obj_val(lp)
	//x1 = glp_get_col_prim(lp, 1);
	x1, err := C.glp_get_col_prim(lp, 1)
	//x2 = glp_get_col_prim(lp, 2);
	x2, err := C.glp_get_col_prim(lp, 2)
	//printf("z = %g; x1 = %g; x2 = %g\n", z, x1, x2);
	fmt.Printf("z = %f; x1 = %f; x2 = %f\n", z, x1, x2)

	/* housekeeping */
	//glp_delete_prob(lp);
	//glp_free_env();
	//return 0;
}
