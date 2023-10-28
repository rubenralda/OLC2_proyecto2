/*------HEADER------*/
#include <stdio.h>
#include <math.h>
double heap[30101999];
double stack[30101999];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34, t35, t36, t37, t38, t39, t40, t41, t42, t43, t44, t45, t46, t47, t48, t49, t50, t51, t52, t53, t54, t55, t56, t57, t58, t59, t60, t61, t62, t63, t64, t65, t66, t67, t68, t69, t70, t71, t72, t73, t74, t75, t76, t77, t78, t79, t80, t81, t82, t83, t84, t85, t86, t87, t88, t89, t90, t91, t92, t93, t94, t95, t96, t97, t98, t99, t100, t101, t102, t103, t104, t105, t106, t107, t108, t109, t110, t111, t112, t113, t114, t115, t116, t117, t118, t119, t120, t121, t122, t123, t124, t125, t126, t127, t128, t129, t130, t131, t132, t133, t134, t135, t136, t137, t138, t139, t140, t141, t142, t143, t144, t145, t146, t147, t148, t149, t150, t151, t152, t153, t154, t155, t156, t157, t158;

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
	heap[(int)H] = 45;
	H = H + 1;
	heap[(int)H] = 65;
	H = H + 1;
	heap[(int)H] = 82;
	H = H + 1;
	heap[(int)H] = 67;
	H = H + 1;
	heap[(int)H] = 72;
	H = H + 1;
	heap[(int)H] = 73;
	H = H + 1;
	heap[(int)H] = 86;
	H = H + 1;
	heap[(int)H] = 79;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 66;
	H = H + 1;
	heap[(int)H] = 65;
	H = H + 1;
	heap[(int)H] = 83;
	H = H + 1;
	heap[(int)H] = 73;
	H = H + 1;
	heap[(int)H] = 67;
	H = H + 1;
	heap[(int)H] = 79;
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
	heap[(int)H] = 49;
	H = H + 1;
	heap[(int)H] = 54;
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
	goto L3;
	L2:
	stack[0] = 1;
	goto L4;
	L3:
	stack[0] = 0;
	goto L4;
	L4:
	
	//Llamando una variable
	t17 = stack[0];
	if(t17 == 1) goto L5;
	goto L6;
	L6:
	stack[1] = 1;
	goto L7;
	L5:
	stack[1] = 0;
	goto L7;
	L7:
	
	t18 = H;
	heap[(int)H] = 105;
	H = H + 1;
	heap[(int)H] = 109;
	H = H + 1;
	heap[(int)H] = 112;
	H = H + 1;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = 105;
	H = H + 1;
	heap[(int)H] = 109;
	H = H + 1;
	heap[(int)H] = 105;
	H = H + 1;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	stack[2] = t18;
	
	t19 = H;
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
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 105;
	H = H + 1;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	stack[3] = t19;
	
	t28 = 2 * 3;
	t27 = 5 + t28;
	t26 = 4 * t27;
	t25 = 2 + t26;
	t24 = 10 * t25;
	t23 = 5 + t24;
	t30 = 8 * 3;
	t29 = t30 * 3;
	t22 = t23 - t29;
	t21 = 7 - t22;
	t32 = 6 * 2;
	t31 = 50 * t32;
	t20 = t21 + t31;
	stack[4] = t20;
	
	t39 = 2 * 2;
	t38 = t39 * 2;
	t37 = t38 * 2;
	t36 = t37 - 9;
	t42 = 8 - 6;
	t47 = 3 * 3;
	t48 = 6 * 5;
	t46 = t47 - t48;
	t45 = t46 - 7;
	t51 = 7 * 7;
	t50 = t51 * 7;
	t49 = 9 + t50;
	t44 = t45 - t49;
	t43 = t44 + 10;
	t41 = t42 + t43;
	t40 = t41 - 5;
	t35 = t36 - t40;
	t34 = t35 + 8;
	t54 = 2 * 3;
	t53 = 5 * t54;
	t52 = 6 - t53;
	t33 = t34 - t52;
	stack[5] = t33;
	
	//Llamando una variable
	t58 = stack[4];
	//Llamando una variable
	t64 = stack[5];
	t62 = t64 * 3;
	t61 = 2 + t62;
	t60 = t61 + 1;
	t68 = 2 * 2;
	t67 = t68 * 2;
	t66 = t67 - 2;
	t65 = t66 * 2;
	t59 = t60 - t65;
	t56 = t58 + t59;
	t55 = t56 - 2;
	stack[6] = t55;
	
	t69 = H;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
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
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 49;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t70 = P + 7;
	t70 = t70 + 1;
	stack[(int)t70] = t69;
	P = P + 7;
	dbrust_printString();
	t71 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	//Llamando una variable
	t73 = stack[4];
	printf("%d", (int)t73);
	printf("%c", 32);
	
	printf("%c", 10);
	t74 = H;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
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
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 50;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t75 = P + 7;
	t75 = t75 + 1;
	stack[(int)t75] = t74;
	P = P + 7;
	dbrust_printString();
	t76 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	//Llamando una variable
	t78 = stack[5];
	printf("%d", (int)t78);
	printf("%c", 32);
	
	printf("%c", 10);
	t79 = H;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
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
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 118;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 51;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t80 = P + 7;
	t80 = t80 + 1;
	stack[(int)t80] = t79;
	P = P + 7;
	dbrust_printString();
	t81 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	//Llamando una variable
	t83 = stack[6];
	printf("%d", (int)t83);
	printf("%c", 32);
	
	printf("%c", 10);
	t84 = H;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = 117;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 116;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = 112;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 114;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 99;
	H = H + 1;
	heap[(int)H] = 105;
	H = H + 1;
	heap[(int)H] = 195;
	H = H + 1;
	heap[(int)H] = 179;
	H = H + 1;
	heap[(int)H] = 110;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t85 = P + 7;
	t85 = t85 + 1;
	stack[(int)t85] = t84;
	P = P + 7;
	dbrust_printString();
	t86 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	//Llamando una variable
	t88 = stack[6];
	printf("%d", (int)t88);
	printf("%c", 32);
	
	printf("%c", 10);
	t89 = H;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
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
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 98;
	H = H + 1;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t90 = P + 7;
	t90 = t90 + 1;
	stack[(int)t90] = t89;
	P = P + 7;
	dbrust_printString();
	t91 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	//Llamando una variable
	t93 = stack[0];
	if(t93 == 1) goto L8;
	goto L9;
	L8:
	printf("%c", (char)116);
	printf("%c", (char)114);
	printf("%c", (char)117);
	printf("%c", (char)101);
	goto L10;
	L9:
	printf("%c", (char)102);
	printf("%c", (char)97);
	printf("%c", (char)108);
	printf("%c", (char)115);
	printf("%c", (char)101);
	L10:
	printf("%c", 32);
	
	printf("%c", 10);
	t94 = H;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
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
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 99;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 49;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t95 = P + 7;
	t95 = t95 + 1;
	stack[(int)t95] = t94;
	P = P + 7;
	dbrust_printString();
	t96 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	//Llamando una variable
	t98 = stack[2];
	t99 = P + 7;
	t99 = t99 + 1;
	stack[(int)t99] = t98;
	P = P + 7;
	dbrust_printString();
	t100 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	printf("%c", 10);
	t101 = H;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
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
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 99;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 50;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t102 = P + 7;
	t102 = t102 + 1;
	stack[(int)t102] = t101;
	P = P + 7;
	dbrust_printString();
	t103 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	//Llamando una variable
	t105 = stack[3];
	t106 = P + 7;
	t106 = t106 + 1;
	stack[(int)t106] = t105;
	P = P + 7;
	dbrust_printString();
	t107 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	printf("%c", 10);
	t108 = H;
	heap[(int)H] = 69;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
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
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 98;
	H = H + 1;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 50;
	H = H + 1;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t109 = P + 7;
	t109 = t109 + 1;
	stack[(int)t109] = t108;
	P = P + 7;
	dbrust_printString();
	t110 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	//Llamando una variable
	t112 = stack[1];
	if(t112 == 1) goto L11;
	goto L12;
	L11:
	printf("%c", (char)116);
	printf("%c", (char)114);
	printf("%c", (char)117);
	printf("%c", (char)101);
	goto L13;
	L12:
	printf("%c", (char)102);
	printf("%c", (char)97);
	printf("%c", (char)108);
	printf("%c", (char)115);
	printf("%c", (char)101);
	L13:
	printf("%c", 32);
	
	printf("%c", 10);
	stack[7] = 100;
	
	stack[8] = 100;
	
	stack[9] = 7;
	
	goto L14;
	L14:
	stack[10] = 1;
	goto L16;
	L15:
	stack[10] = 0;
	goto L16;
	L16:
	
	stack[11] = 10.0;
	
	stack[12] = 10.0;
	
	//Llamando una variable
	t116 = stack[7];
	//Llamando una variable
	t118 = stack[8];
	if(t116 > t118) goto L17;
	goto L18;
	L18:
	//Llamando una variable
	t121 = stack[8];
	//Llamando una variable
	t123 = stack[9];
	if(t121 < t123) goto L19;
	goto L20;
	L19:
	L17:
	printf("%c", (char)116);
	printf("%c", (char)114);
	printf("%c", (char)117);
	printf("%c", (char)101);
	goto L21;
	L20:
	printf("%c", (char)102);
	printf("%c", (char)97);
	printf("%c", (char)108);
	printf("%c", (char)115);
	printf("%c", (char)101);
	L21:
	printf("%c", 32);
	
	printf("%c", 10);
	//Llamando una variable
	t128 = stack[7];
	//Llamando una variable
	t130 = stack[8];
	if(t128 == t130) goto L22;
	goto L23;
	L22:
	//Llamando una variable
	t133 = stack[11];
	//Llamando una variable
	t135 = stack[12];
	if(t133 == t135) goto L24;
	goto L25;
	L25:
	L23:
	//Llamando una variable
	t138 = stack[9];
	if(14 != t138) goto L26;
	goto L27;
	L26:
	L24:
	printf("%c", (char)116);
	printf("%c", (char)114);
	printf("%c", (char)117);
	printf("%c", (char)101);
	goto L28;
	L27:
	printf("%c", (char)102);
	printf("%c", (char)97);
	printf("%c", (char)108);
	printf("%c", (char)115);
	printf("%c", (char)101);
	L28:
	printf("%c", 32);
	
	printf("%c", 10);
	stack[13] = 5;
	
	stack[14] = 5;
	
	stack[15] = 100;
	
	//Llamando una variable
	t142 = stack[15];
	t144 = 50 + 50;
	//Llamando una variable
	t147 = stack[13];
	//Llamando una variable
	t149 = stack[13];
	t145 = t147 - t149;
	t143 = t144 + t145;
	if(t142 == t143) goto L29;
	goto L30;
	L29:
	goto L31;
	L31:
	printf("%c", (char)116);
	printf("%c", (char)114);
	printf("%c", (char)117);
	printf("%c", (char)101);
	goto L33;
	L32:
	L30:
	printf("%c", (char)102);
	printf("%c", (char)97);
	printf("%c", (char)108);
	printf("%c", (char)115);
	printf("%c", (char)101);
	L33:
	printf("%c", 32);
	
	printf("%c", 10);
	stack[16] = 15;
	
	//Llamando una variable
	t155 = stack[16];
	t153 = (int)t155 % (int)2;
	if(t153 == 0) goto L34;
	goto L35;
	L34:
	printf("%c", (char)116);
	printf("%c", (char)114);
	printf("%c", (char)117);
	printf("%c", (char)101);
	goto L36;
	L35:
	printf("%c", (char)102);
	printf("%c", (char)97);
	printf("%c", (char)108);
	printf("%c", (char)115);
	printf("%c", (char)101);
	L36:
	printf("%c", 32);
	
	printf("%c", 10);
	t156 = H;
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
	heap[(int)H] = 110;
	H = H + 1;
	heap[(int)H] = 117;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	
	t157 = P + 17;
	t157 = t157 + 1;
	stack[(int)t157] = t156;
	P = P + 17;
	dbrust_printString();
	t158 = stack[(int)P];
	P = P - 17;
	printf("%c", 32);
	
	printf("%c", (char)110);
	printf("%c", (char)105);
	printf("%c", (char)108);
	printf("%c", 32);
	
	printf("%c", 10);

	return 0;
}
