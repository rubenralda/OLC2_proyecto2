/*------HEADER------*/
#include <stdio.h>
#include <math.h>
double heap[30101999];
double stack[30101999];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34, t35, t36, t37, t38, t39, t40, t41, t42, t43, t44, t45, t46, t47, t48, t49, t50, t51, t52, t53, t54, t55, t56, t57, t58, t59, t60, t61, t62, t63, t64, t65, t66, t67, t68, t69, t70, t71, t72, t73, t74, t75, t76, t77, t78, t79, t80, t81, t82, t83, t84, t85, t86, t87, t88, t89, t90, t91, t92, t93, t94, t95, t96, t97, t98, t99, t100, t101, t102, t103, t104, t105, t106, t107, t108, t109, t110, t111, t112, t113, t114, t115, t116, t117, t118, t119, t120, t121, t122, t123, t124, t125, t126, t127, t128, t129, t130, t131, t132, t133, t134, t135, t136, t137, t138, t139, t140, t141, t142, t143, t144, t145, t146, t147, t148, t149, t150, t151, t152, t153, t154, t155, t156, t157, t158, t159, t160, t161, t162, t163, t164, t165, t166, t167, t168, t169, t170, t171, t172, t173, t174, t175, t176, t177, t178, t179, t180, t181, t182, t183, t184, t185, t186, t187, t188, t189, t190, t191, t192, t193, t194, t195, t196, t197, t198, t199, t200, t201, t202, t203, t204, t205;

/*------NATIVES------*/
void dbrust_printString() {
	t1 = P + 1;
	t2 = stack[(int)t1];
	L1:
	t3 = heap[(int)t2];
	if(t3 == -1) goto L0;
	printf("%c", (char)t3);
	t2 = t2 + 1;
	goto L1;
	L0:
	return;
}

void funcion_float_rubencin() {
	t104 = P + 1;
	t106 = stack[(int)t104];
	t111 = t106;
	L11:
	t105 = heap[(int)t106];
	if(t105 == -1) goto L14;
	if(t105 == 46) goto L13;
	t106 = t106 + 1;
	goto L11;
	L13:
	t109 = t106 + 1;
	goto L12;
	L14:
	t109 = t106;
	L12:
	t106 = t106 - 1;
	t108 = 1;
	t107 = 0;
	L9:
	if(t106 < t111) goto L10;
	t105 = heap[(int)t106];
	if(t105 < 48) goto L15;
	if(t105 > 57) goto L15;
	t106 = t106 - 1;
	t105 = t105 - 48;
	t105 = t105 * t108;
	t108 = t108 * 10;
	t107 = t107 + t105;
	goto L9;
	L10:
	t108 = 10;
	t110:
	t105 = heap[(int)t109];
	if(t105 == -1) goto L8;
	if(t105 < 48) goto L15;
	if(t105 > 57) goto L15;
	t109 = t109 + 1;
	t105 = t105 - 48;
	t105 = t105 / t108;
	t108 = t108 * 10;
	t107 = t107 + t105;
	goto t110;
	L15:
	printf("%c", (char) 69);
	printf("%c", (char) 82);
	printf("%c", (char) 82);
	printf("%c", (char) 79);
	printf("%c", (char) 82);
	t107 = 0;
	L8:
	stack[(int)P] = t107;
	return;
}

void funcion_int_float_rubencin() {
	t120 = P + 1;
	t121 = stack[(int)t120];
	t121 = (int)t121;
	stack[(int)P] = t121;
	return;
}

void funcion_int_string_rubencin() {
	t120 = P + 1;
	t123 = stack[(int)t120];
	L19:
	t122 = heap[(int)t123];
	if(t122 == -1) goto L20;
	if(t122 == 46) goto L20;
	t123 = t123 + 1;
	goto L19;
	L20:
	t123 = t123 - 1;
	t125 = 1;
	t124 = 0;
	L17:
	t122 = heap[(int)t123];
	if(t122 == -1) goto L16;
	t123 = t123 - 1;
	if(t122 < 48) goto L18;
	if(t122 > 57) goto L18;
	t122 = t122 - 48;
	t122 = t122 * t125;
	t125 = t125 * 10;
	t124 = t124 + t122;
	goto L17;
	L18:
	if(t122 == 46) goto L16;
	printf("%c", (char) 69);
	printf("%c", (char) 82);
	printf("%c", (char) 82);
	printf("%c", (char) 79);
	printf("%c", (char) 82);
	t124 = 0;
	L16:
	stack[(int)P] = t124;
	return;
}

void funcion_string_rubencin() {
	t140 = P + 1;
	t141 = stack[(int)t140];
	stack[(int)P] = H;
	t143 = t141 - (int)t141;
	t142 = (int)t141;
	t145 = 0;
	L21:
	if(t142 < 1) goto L22;
	t142 = t142/ 10;
	t145 = t145 + 1;
	goto L21;
	L22:
	t142 = t142* 10;
	t144 = (int)t142;
	t144 = t144 + 48;
	heap[(int)H] = t144;
	H = H + 1;
	t142 = t142 - (int)t142;
	if((int)t145 <= 1) goto L24;
	t145 = t145 - 1;
	goto L22;
	L24:
	if(t143 == 0) goto L25;
	heap[(int)H] = 46;
	H = H + 1;
	t146 = 6;
	L23:
	t143 = t143* 10;
	t144 = (int)t143;
	t144 = t144 + 48;
	heap[(int)H] = t144;
	H = H + 1;
	t143 = t143 - (int)t143;
	if((int)t146 <= 1) goto L25;
	t146 = t146 - 1;
	goto L23;
	L25:
	heap[(int)H] = -1;
	H = H + 1;
	return;
}

/*------FUNCTIONS------*/
void suma() {
t21 = P - 0;
t19 = t21 + 1;
t20 = stack[(int)t19];
t24 = P - 0;
t22 = t24 + 2;
t23 = stack[(int)t22];
t25 = t20 + t23;
t25 = (int)t25;
t18 = P + 3;
stack[(int)t18] = t25;

//bananero
t28 = P - 0;
t26 = t28 + 3;
t27 = stack[(int)t26];
P = P - 0;
stack[(int)P] = t27;
goto L2;
L2:
	return;
}

void saludo3() {
t41 = H;
heap[(int)H] = 115;
H = H + 1;
heap[(int)H] = 97;
H = H + 1;
heap[(int)H] = 108;
H = H + 1;
heap[(int)H] = 117;
H = H + 1;
heap[(int)H] = 100;
H = H + 1;
heap[(int)H] = 111;
H = H + 1;
heap[(int)H] = 115;
H = H + 1;
heap[(int)H] = 33;
H = H + 1;
heap[(int)H] = -1;
H = H + 1;

t42 = P + 1;
t42 = t42 + 1;
stack[(int)t42] = t41;
P = P + 1;
dbrust_printString();
t43 = stack[(int)P];
P = P - 1;
printf("%c", 32);

printf("%c", 10);
L3:
	return;
}

void saludo2() {
t44 = H;
heap[(int)H] = 109;
H = H + 1;
heap[(int)H] = 117;
H = H + 1;
heap[(int)H] = 110;
H = H + 1;
heap[(int)H] = 100;
H = H + 1;
heap[(int)H] = 111;
H = H + 1;
heap[(int)H] = -1;
H = H + 1;

t45 = P + 1;
t45 = t45 + 1;
stack[(int)t45] = t44;
P = P + 1;
dbrust_printString();
t46 = stack[(int)P];
P = P - 1;
printf("%c", 32);

printf("%c", 10);
//llamadaempieza
t47 = P + 2;
P = P + 1;
saludo3();
t48 = stack[(int)P];
P = P - 1;
//finLlamada
L4:
	return;
}

void saludo1() {
t49 = H;
heap[(int)H] = 104;
H = H + 1;
heap[(int)H] = 111;
H = H + 1;
heap[(int)H] = 108;
H = H + 1;
heap[(int)H] = 97;
H = H + 1;
heap[(int)H] = -1;
H = H + 1;

t50 = P + 1;
t50 = t50 + 1;
stack[(int)t50] = t49;
P = P + 1;
dbrust_printString();
t51 = stack[(int)P];
P = P - 1;
printf("%c", 32);

printf("%c", 10);
//llamadaempieza
t52 = P + 2;
P = P + 1;
saludo2();
t53 = stack[(int)P];
P = P - 1;
//finLlamada
L5:
	return;
}

void ejemplo2() {
t61 = P - 0;
t59 = t61 + 1;
t60 = stack[(int)t59];
printf("%d", (int)t60);
printf("%c", 32);
printf("%c", 10);
t64 = P - 0;
t62 = t64 + 2;
t63 = stack[(int)t62];
printf("%d", (int)t63);
printf("%c", 32);
printf("%c", 10);
L6:
	return;
}

void duplicar() {
t80 = P - 0;
t78 = t80 + 1;
t81 = stack[(int)t78];
t79 = stack[(int)t81];
t83 = P + 0;
t82 = t83 + 1;
t86 = stack[(int)t82];
t84 = stack[(int)t86];
t85 = t84 + t79;
stack[(int)t86] = t85;
L7:
	return;
}

/*------MAIN------*/
int main() {
	P = 0; H = 0;

	t0 = H;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t4 = P + 0;
	t4 = t4 + 1;
	stack[(int)t4] = t0;
	P = P + 0;
	dbrust_printString();
	t5 = stack[(int)P];
	P = P - 0;
	printf("%c", 32);
	
	printf("%c", 10);
	t6 = H;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 70;
	H = H + 1;
	heap[(int)H] = 85;
	H = H + 1;
	heap[(int)H] = 78;
	H = H + 1;
	heap[(int)H] = 67;
	H = H + 1;
	heap[(int)H] = 73;
	H = H + 1;
	heap[(int)H] = 79;
	H = H + 1;
	heap[(int)H] = 78;
	H = H + 1;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = 83;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = 77;
	H = H + 1;
	heap[(int)H] = 66;
	H = H + 1;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = 66;
	H = H + 1;
	heap[(int)H] = 73;
	H = H + 1;
	heap[(int)H] = 68;
	H = H + 1;
	heap[(int)H] = 65;
	H = H + 1;
	heap[(int)H] = 83;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t7 = P + 0;
	t7 = t7 + 1;
	stack[(int)t7] = t6;
	P = P + 0;
	dbrust_printString();
	t8 = stack[(int)P];
	P = P - 0;
	printf("%c", 32);
	
	printf("%c", 10);
	t9 = H;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 49;
	H = H + 1;
	heap[(int)H] = 52;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 112;
	H = H + 1;
	heap[(int)H] = 116;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t10 = P + 0;
	t10 = t10 + 1;
	stack[(int)t10] = t9;
	P = P + 0;
	dbrust_printString();
	t11 = stack[(int)P];
	P = P - 0;
	printf("%c", 32);
	
	printf("%c", 10);
	t12 = H;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t13 = P + 0;
	t13 = t13 + 1;
	stack[(int)t13] = t12;
	P = P + 0;
	dbrust_printString();
	t14 = stack[(int)P];
	P = P - 0;
	printf("%c", 32);
	
	printf("%c", 10);
	t15 = H;
	heap[(int)H] = -1;
	H = H + 1;
	
	t16 = P + 0;
	t16 = t16 + 1;
	stack[(int)t16] = t15;
	P = P + 0;
	dbrust_printString();
	t17 = stack[(int)P];
	P = P - 0;
	printf("%c", 32);
	
	printf("%c", 10);
	//llamadaempieza
	t30 = P + 1;
	stack[(int)t30] = 5;
	t30 = t30 + 1;
	stack[(int)t30] = 3;
	t30 = t30 + 1;
	P = P + 0;
	suma();
	t31 = stack[(int)P];
	P = P - 0;
	//finLlamada
	t29 = P + 0;
	stack[(int)t29] = t31;
	
	t32 = H;
	heap[(int)H] = 76;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = 117;
	H = H + 1;
	heap[(int)H] = 109;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t33 = P + 1;
	t33 = t33 + 1;
	stack[(int)t33] = t32;
	P = P + 1;
	dbrust_printString();
	t34 = stack[(int)P];
	P = P - 1;
	printf("%c", 32);
	
	t37 = P - 0;
	t35 = t37 + 0;
	t36 = stack[(int)t35];
	printf("%d", (int)t36);
	printf("%c", 32);
	printf("%c", 10);
	t38 = H;
	heap[(int)H] = -1;
	H = H + 1;
	
	t39 = P + 1;
	t39 = t39 + 1;
	stack[(int)t39] = t38;
	P = P + 1;
	dbrust_printString();
	t40 = stack[(int)P];
	P = P - 1;
	printf("%c", 32);
	
	printf("%c", 10);
	//llamadaempieza
	t54 = P + 2;
	P = P + 1;
	saludo1();
	t55 = stack[(int)P];
	P = P - 1;
	//finLlamada
	t56 = H;
	heap[(int)H] = -1;
	H = H + 1;
	
	t57 = P + 1;
	t57 = t57 + 1;
	stack[(int)t57] = t56;
	P = P + 1;
	dbrust_printString();
	t58 = stack[(int)P];
	P = P - 1;
	printf("%c", 32);
	
	printf("%c", 10);
	t65 = P + 1;
	stack[(int)t65] = 251;
	
	t66 = P + 2;
	stack[(int)t66] = 85;
	
	//llamadaempieza
	t67 = P + 4;
	t70 = P - 0;
	t68 = t70 + 1;
	t69 = stack[(int)t68];
	stack[(int)t67] = t69;
	t67 = t67 + 1;
	t73 = P - 0;
	t71 = t73 + 2;
	t72 = stack[(int)t71];
	stack[(int)t67] = t72;
	t67 = t67 + 1;
	P = P + 3;
	ejemplo2();
	t74 = stack[(int)P];
	P = P - 3;
	//finLlamada
	t75 = H;
	heap[(int)H] = -1;
	H = H + 1;
	
	t76 = P + 3;
	t76 = t76 + 1;
	stack[(int)t76] = t75;
	P = P + 3;
	dbrust_printString();
	t77 = stack[(int)P];
	P = P - 3;
	printf("%c", 32);
	
	printf("%c", 10);
	t87 = P + 3;
	stack[(int)t87] = 1;
	
	//llamadaempieza
	t88 = P + 5;
	t91 = P - 0;
	t89 = t91 + 3;
	t90 = stack[(int)t89];
	stack[(int)t88] = t89;
	t88 = t88 + 1;
	P = P + 4;
	duplicar();
	t92 = stack[(int)P];
	P = P - 4;
	//finLlamada
	t93 = H;
	heap[(int)H] = 110;
	H = H + 1;
	heap[(int)H] = 117;
	H = H + 1;
	heap[(int)H] = 109;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = 50;
	H = H + 1;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t94 = P + 4;
	t94 = t94 + 1;
	stack[(int)t94] = t93;
	P = P + 4;
	dbrust_printString();
	t95 = stack[(int)P];
	P = P - 4;
	printf("%c", 32);
	
	t98 = P - 0;
	t96 = t98 + 3;
	t97 = stack[(int)t96];
	printf("%d", (int)t97);
	printf("%c", 32);
	printf("%c", 10);
	t99 = H;
	heap[(int)H] = -1;
	H = H + 1;
	
	t100 = P + 4;
	t100 = t100 + 1;
	stack[(int)t100] = t99;
	P = P + 4;
	dbrust_printString();
	t101 = stack[(int)P];
	P = P - 4;
	printf("%c", 32);
	
	printf("%c", 10);
	t103 = H;
	heap[(int)H] = 57;
	H = H + 1;
	heap[(int)H] = 46;
	H = H + 1;
	heap[(int)H] = 53;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t112 = P + 4;
	t112 = t112 + 1;
	stack[(int)t112] = t103;
	P = P + 4;
	funcion_float_rubencin();
	t113 = stack[(int)P];
	P = P - 4;
	t102 = P + 4;
	stack[(int)t102] = t113;
	
	t115 = H;
	heap[(int)H] = 51;
	H = H + 1;
	heap[(int)H] = 46;
	H = H + 1;
	heap[(int)H] = 54;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t116 = P + 5;
	t116 = t116 + 1;
	stack[(int)t116] = t115;
	P = P + 5;
	funcion_float_rubencin();
	t117 = stack[(int)P];
	P = P - 5;
	t114 = P + 5;
	stack[(int)t114] = t117;
	
	t119 = H;
	heap[(int)H] = 50;
	H = H + 1;
	heap[(int)H] = 48;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t126 = P + 6;
	t126 = t126 + 1;
	stack[(int)t126] = t119;
	P = P + 6;
	funcion_int_string_rubencin();
	t127 = stack[(int)P];
	P = P - 6;
	t118 = P + 6;
	stack[(int)t118] = t127;
	
	t129 = H;
	heap[(int)H] = 50;
	H = H + 1;
	heap[(int)H] = 48;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t130 = P + 7;
	t130 = t130 + 1;
	stack[(int)t130] = t129;
	P = P + 7;
	funcion_int_string_rubencin();
	t131 = stack[(int)P];
	P = P - 7;
	t128 = P + 7;
	stack[(int)t128] = t131;
	
	t135 = P - 0;
	t133 = t135 + 6;
	t134 = stack[(int)t133];
	t138 = P - 0;
	t136 = t138 + 7;
	t137 = stack[(int)t136];
	t139 = t134 + t137;
	t139 = (int)t139;
	t147 = P + 8;
	t147 = t147 + 1;
	stack[(int)t147] = t139;
	P = P + 8;
	funcion_string_rubencin();
	t148 = stack[(int)P];
	P = P - 8;
	t132 = P + 8;
	stack[(int)t132] = t148;
	
	t152 = P - 0;
	t150 = t152 + 4;
	t151 = stack[(int)t150];
	t155 = P - 0;
	t153 = t155 + 5;
	t154 = stack[(int)t153];
	t156 = t151 - t154;
	t157 = P + 9;
	t157 = t157 + 1;
	stack[(int)t157] = t156;
	P = P + 9;
	funcion_string_rubencin();
	t158 = stack[(int)P];
	P = P - 9;
	t149 = P + 9;
	stack[(int)t149] = t158;
	
	goto L26;
	//momoo
	L26:
	t160 = H;
	heap[(int)H] = 116;
	H = H + 1;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = 117;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	goto L28;
	L27:
	t160 = H;
	heap[(int)H] = 102;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	L28:
	t161 = H;
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 62;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t163 = H;
	L30:
	t164 = heap[(int)t160];
	t160 = t160 + 1;
	if(t164 == -1) goto L29;
	heap[(int)H] = t164;
	H = H + 1;
	goto L30;
	L29:
	t164 = heap[(int)t161];
	t161 = t161 + 1;
	if(t164 == -1) goto L31;
	heap[(int)H] = t164;
	H = H + 1;
	goto L29;
	L31:
	heap[(int)H] = -1;
	H = H + 1;
	
	t165 = P + 10;
	t165 = t165 + 1;
	stack[(int)t165] = 3.504;
	P = P + 10;
	funcion_string_rubencin();
	t166 = stack[(int)P];
	P = P - 10;
	t168 = H;
	L33:
	t169 = heap[(int)t163];
	t163 = t163 + 1;
	if(t169 == -1) goto L32;
	heap[(int)H] = t169;
	H = H + 1;
	goto L33;
	L32:
	t169 = heap[(int)t166];
	t166 = t166 + 1;
	if(t169 == -1) goto L34;
	heap[(int)H] = t169;
	H = H + 1;
	goto L32;
	L34:
	heap[(int)H] = -1;
	H = H + 1;
	
	t159 = P + 10;
	stack[(int)t159] = t168;
	
	t170 = H;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 116;
	H = H + 1;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = 117;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t171 = P + 11;
	t171 = t171 + 1;
	stack[(int)t171] = t170;
	P = P + 11;
	dbrust_printString();
	t172 = stack[(int)P];
	P = P - 11;
	printf("%c", 32);
	
	goto L35;
	//momoo
	L35:
	t173 = H;
	heap[(int)H] = 116;
	H = H + 1;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = 117;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	goto L37;
	L36:
	t173 = H;
	heap[(int)H] = 102;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	L37:
	t174 = P + 11;
	t174 = t174 + 1;
	stack[(int)t174] = t173;
	P = P + 11;
	dbrust_printString();
	t175 = stack[(int)P];
	P = P - 11;
	printf("%c", 32);
	
	printf("%c", 10);
	t176 = H;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 102;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t177 = P + 11;
	t177 = t177 + 1;
	stack[(int)t177] = t176;
	P = P + 11;
	dbrust_printString();
	t178 = stack[(int)P];
	P = P - 11;
	printf("%c", 32);
	
	goto L39;
	//momoo
	L38:
	t179 = H;
	heap[(int)H] = 116;
	H = H + 1;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = 117;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	goto L40;
	L39:
	t179 = H;
	heap[(int)H] = 102;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	L40:
	t180 = P + 11;
	t180 = t180 + 1;
	stack[(int)t180] = t179;
	P = P + 11;
	dbrust_printString();
	t181 = stack[(int)P];
	P = P - 11;
	printf("%c", 32);
	
	printf("%c", 10);
	t182 = H;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 49;
	H = H + 1;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t183 = P + 11;
	t183 = t183 + 1;
	stack[(int)t183] = t182;
	P = P + 11;
	dbrust_printString();
	t184 = stack[(int)P];
	P = P - 11;
	printf("%c", 32);
	
	t187 = P - 0;
	t185 = t187 + 8;
	t186 = stack[(int)t185];
	t188 = P + 11;
	t188 = t188 + 1;
	stack[(int)t188] = t186;
	P = P + 11;
	dbrust_printString();
	t189 = stack[(int)P];
	P = P - 11;
	printf("%c", 32);
	
	printf("%c", 10);
	t190 = H;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 50;
	H = H + 1;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t191 = P + 11;
	t191 = t191 + 1;
	stack[(int)t191] = t190;
	P = P + 11;
	dbrust_printString();
	t192 = stack[(int)P];
	P = P - 11;
	printf("%c", 32);
	
	t195 = P - 0;
	t193 = t195 + 9;
	t194 = stack[(int)t193];
	t196 = P + 11;
	t196 = t196 + 1;
	stack[(int)t196] = t194;
	P = P + 11;
	dbrust_printString();
	t197 = stack[(int)P];
	P = P - 11;
	printf("%c", 32);
	
	printf("%c", 10);
	t198 = H;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 99;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 110;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t199 = P + 11;
	t199 = t199 + 1;
	stack[(int)t199] = t198;
	P = P + 11;
	dbrust_printString();
	t200 = stack[(int)P];
	P = P - 11;
	printf("%c", 32);
	
	t203 = P - 0;
	t201 = t203 + 10;
	t202 = stack[(int)t201];
	t204 = P + 11;
	t204 = t204 + 1;
	stack[(int)t204] = t202;
	P = P + 11;
	dbrust_printString();
	t205 = stack[(int)P];
	P = P - 11;
	printf("%c", 32);
	
	printf("%c", 10);

	return 0;
}
