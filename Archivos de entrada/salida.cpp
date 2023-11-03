/*------HEADER------*/
#include <stdio.h>
#include <math.h>
double heap[30101999];
double stack[30101999];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26;

/*------NATIVES------*/
void funcion_float_rubencin() {
	t6 = P + 1;
	t8 = stack[(int)t6];
	t13 = t8;
	L3:
	t7 = heap[(int)t8];
	if(t7 == -1) goto L6;
	if(t7 == 46) goto L5;
	t8 = t8 + 1;
	goto L3;
	L5:
	t11 = t8 + 1;
	goto L4;
	L6:
	t11 = t8;
	L4:
	t8 = t8 - 1;
	t10 = 1;
	t9 = 0;
	L1:
	if(t8 < t13) goto L2;
	t7 = heap[(int)t8];
	if(t7 < 48) goto L7;
	if(t7 > 57) goto L7;
	t8 = t8 - 1;
	t7 = t7 - 48;
	t7 = t7 * t10;
	t10 = t10 * 10;
	t9 = t9 + t7;
	goto L1;
	L2:
	t10 = 10;
	t12:
	t7 = heap[(int)t11];
	if(t7 == -1) goto L0;
	if(t7 < 48) goto L7;
	if(t7 > 57) goto L7;
	t11 = t11 + 1;
	t7 = t7 - 48;
	t7 = t7 / t10;
	t10 = t10 * 10;
	t9 = t9 + t7;
	goto t12;
	L7:
	printf("%c", (char) 69);
	printf("%c", (char) 82);
	printf("%c", (char) 82);
	printf("%c", (char) 79);
	printf("%c", (char) 82);
	t9 = 0;
	L0:
	stack[(int)P] = t9;
	return;
}

/*------MAIN------*/
int main() {
	P = 0; H = 0;

	//inicio declaracion
	t2 = P + 0;
	stack[(int)t2] = H;
	t0 = H;
	t1 = t0 + 1;
	heap[(int)t1] = -1;
	H = H + 2;
	//ejecuta expresion
	t3 = H;
	H = H + 1;
	t5 = H;
	heap[(int)H] = 50;
	H = H + 1;
	heap[(int)H] = 51;
	H = H + 1;
	heap[(int)H] = 46;
	H = H + 1;
	heap[(int)H] = 52;
	H = H + 1;
	heap[(int)H] = 51;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t14 = P + 1;
	t14 = t14 + 1;
	stack[(int)t14] = t5;
	P = P + 1;
	funcion_float_rubencin();
	t15 = stack[(int)P];
	P = P - 1;
	t4 = t3 + 0;
	heap[(int)t4] = t15;
	H = H + 1;
	//metiendovalor
	heap[(int)t0] = t3;
	t1 = t0 + 1;
	heap[(int)t1] = H;
	t0 = H;
	t1 = t0 + 1;
	heap[(int)t1] = -1;
	H = H + 2;
	//metiendovalortermino
	//ejecuta expresion
	t16 = H;
	H = H + 1;
	t18 = H;
	heap[(int)H] = 57;
	H = H + 1;
	heap[(int)H] = 54;
	H = H + 1;
	heap[(int)H] = 46;
	H = H + 1;
	heap[(int)H] = 51;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t19 = P + 1;
	t19 = t19 + 1;
	stack[(int)t19] = t18;
	P = P + 1;
	funcion_float_rubencin();
	t20 = stack[(int)P];
	P = P - 1;
	t17 = t16 + 0;
	heap[(int)t17] = t20;
	H = H + 1;
	//metiendovalor
	heap[(int)t0] = t16;
	t1 = t0 + 1;
	heap[(int)t1] = H;
	t0 = H;
	t1 = t0 + 1;
	heap[(int)t1] = -1;
	H = H + 2;
	//metiendovalortermino
	t22 = P - 0;
	t21 = t22 + 0;
	t23 = stack[(int)t21];
	t25 = 1;
	//metiendo
	L8:
	t26 = heap[(int)t23];
	t23 = t23 + 1;
	t23 = heap[(int)t23];
	if(t23 == -1) goto L9;
	if(t25 <= 0) goto L10;
	t25 = t25 - 1;
	goto L8;
	L9:
	printf("%c", 69);
	printf("%c", 114);
	printf("%c", 114);
	printf("%c", 111);
	printf("%c", 114);
	t26 = 0;
	L10:
	t23 = t26;
	//termino
	t23 = t23 + 0;
	t24 = heap[(int)t23];
	printf("%f", t24);
	printf("%c", 32);
	printf("%c", 10);

	return 0;
}
