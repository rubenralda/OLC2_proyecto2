/*------HEADER------*/
#include <stdio.h>
#include <math.h>
double heap[30101999];
double stack[30101999];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34, t35, t36, t37, t38, t39, t40, t41, t42, t43, t44, t45, t46, t47, t48, t49, t50, t51, t52, t53, t54, t55, t56, t57, t58, t59, t60, t61, t62, t63, t64, t65, t66, t67, t68, t69, t70, t71, t72, t73, t74, t75, t76, t77, t78, t79, t80, t81, t82, t83, t84, t85, t86, t87, t88, t89, t90, t91, t92, t93, t94, t95, t96, t97, t98, t99, t100, t101, t102, t103, t104, t105, t106, t107, t108, t109, t110, t111, t112, t113, t114, t115, t116, t117, t118, t119, t120, t121, t122, t123, t124, t125, t126, t127, t128, t129, t130, t131, t132, t133, t134, t135, t136, t137, t138, t139, t140, t141, t142, t143, t144, t145, t146, t147, t148, t149, t150, t151, t152, t153, t154, t155, t156, t157, t158, t159, t160, t161, t162, t163, t164, t165, t166, t167, t168, t169, t170, t171, t172, t173, t174, t175, t176, t177, t178, t179, t180, t181, t182, t183, t184, t185;

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
	t15 = P + 0;
	stack[(int)t15] = 1;
	goto L4;
	L3:
	t15 = P + 0;
	stack[(int)t15] = 0;
	goto L4;
	L4:
	
	t19 = P - 0;
	t17 = t19 + 0;
	t18 = stack[(int)t17];
	if(t18 == 1) goto L5;
	goto L6;
	L6:
	t16 = P + 1;
	stack[(int)t16] = 1;
	goto L7;
	L5:
	t16 = P + 1;
	stack[(int)t16] = 0;
	goto L7;
	L7:
	
	t21 = H;
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
	
	t20 = P + 2;
	stack[(int)t20] = t21;
	
	t23 = H;
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
	
	t22 = P + 3;
	stack[(int)t22] = t23;
	
	t25 = 2 * 3;
	t26 = 5 + t25;
	t27 = 4 * t26;
	t28 = 2 + t27;
	t29 = 10 * t28;
	t30 = 5 + t29;
	t31 = 8 * 3;
	t32 = t31 * 3;
	t33 = t30 - t32;
	t34 = 7 - t33;
	t35 = 6 * 2;
	t36 = 50 * t35;
	t37 = t34 + t36;
	t24 = P + 4;
	stack[(int)t24] = t37;
	
	t39 = 2 * 2;
	t40 = t39 * 2;
	t41 = t40 * 2;
	t42 = t41 - 9;
	t43 = 8 - 6;
	t44 = 3 * 3;
	t45 = 6 * 5;
	t46 = t44 - t45;
	t47 = t46 - 7;
	t48 = 7 * 7;
	t49 = t48 * 7;
	t50 = 9 + t49;
	t51 = t47 - t50;
	t52 = t51 + 10;
	t53 = t43 + t52;
	t54 = t53 - 5;
	t55 = t42 - t54;
	t56 = t55 + 8;
	t57 = 2 * 3;
	t58 = 5 * t57;
	t59 = 6 - t58;
	t60 = t56 - t59;
	t38 = P + 5;
	stack[(int)t38] = t60;
	
	t64 = P - 0;
	t62 = t64 + 4;
	t63 = stack[(int)t62];
	t67 = P - 0;
	t65 = t67 + 5;
	t66 = stack[(int)t65];
	t68 = t66 * 3;
	t69 = 2 + t68;
	t70 = t69 + 1;
	t71 = 2 * 2;
	t72 = t71 * 2;
	t73 = t72 - 2;
	t74 = t73 * 2;
	t75 = t70 - t74;
	t76 = t63 + t75;
	t77 = t76 - 2;
	t61 = P + 6;
	stack[(int)t61] = t77;
	
	t78 = H;
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
	
	t79 = P + 7;
	t79 = t79 + 1;
	stack[(int)t79] = t78;
	P = P + 7;
	dbrust_printString();
	t80 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	t83 = P - 0;
	t81 = t83 + 4;
	t82 = stack[(int)t81];
	printf("%d", (int)t82);
	printf("%c", 32);
	
	printf("%c", 10);
	t84 = H;
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
	
	t85 = P + 7;
	t85 = t85 + 1;
	stack[(int)t85] = t84;
	P = P + 7;
	dbrust_printString();
	t86 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	t89 = P - 0;
	t87 = t89 + 5;
	t88 = stack[(int)t87];
	printf("%d", (int)t88);
	printf("%c", 32);
	
	printf("%c", 10);
	t90 = H;
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
	
	t91 = P + 7;
	t91 = t91 + 1;
	stack[(int)t91] = t90;
	P = P + 7;
	dbrust_printString();
	t92 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	t95 = P - 0;
	t93 = t95 + 6;
	t94 = stack[(int)t93];
	printf("%d", (int)t94);
	printf("%c", 32);
	
	printf("%c", 10);
	t96 = H;
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
	
	t97 = P + 7;
	t97 = t97 + 1;
	stack[(int)t97] = t96;
	P = P + 7;
	dbrust_printString();
	t98 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	t101 = P - 0;
	t99 = t101 + 6;
	t100 = stack[(int)t99];
	printf("%d", (int)t100);
	printf("%c", 32);
	
	printf("%c", 10);
	t102 = H;
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
	
	t103 = P + 7;
	t103 = t103 + 1;
	stack[(int)t103] = t102;
	P = P + 7;
	dbrust_printString();
	t104 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	t107 = P - 0;
	t105 = t107 + 0;
	t106 = stack[(int)t105];
	if(t106 == 1) goto L8;
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
	
	t109 = P + 7;
	t109 = t109 + 1;
	stack[(int)t109] = t108;
	P = P + 7;
	dbrust_printString();
	t110 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	t113 = P - 0;
	t111 = t113 + 2;
	t112 = stack[(int)t111];
	t114 = P + 7;
	t114 = t114 + 1;
	stack[(int)t114] = t112;
	P = P + 7;
	dbrust_printString();
	t115 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	printf("%c", 10);
	t116 = H;
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
	
	t117 = P + 7;
	t117 = t117 + 1;
	stack[(int)t117] = t116;
	P = P + 7;
	dbrust_printString();
	t118 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	t121 = P - 0;
	t119 = t121 + 3;
	t120 = stack[(int)t119];
	t122 = P + 7;
	t122 = t122 + 1;
	stack[(int)t122] = t120;
	P = P + 7;
	dbrust_printString();
	t123 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	printf("%c", 10);
	t124 = H;
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
	
	t125 = P + 7;
	t125 = t125 + 1;
	stack[(int)t125] = t124;
	P = P + 7;
	dbrust_printString();
	t126 = stack[(int)P];
	P = P - 7;
	printf("%c", 32);
	
	t129 = P - 0;
	t127 = t129 + 1;
	t128 = stack[(int)t127];
	if(t128 == 1) goto L11;
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
	t130 = P + 7;
	stack[(int)t130] = 100;
	
	t131 = P + 8;
	stack[(int)t131] = 100;
	
	t132 = P + 9;
	stack[(int)t132] = 7;
	
	goto L14;
	L14:
	t133 = P + 10;
	stack[(int)t133] = 1;
	goto L16;
	L15:
	t133 = P + 10;
	stack[(int)t133] = 0;
	goto L16;
	L16:
	
	t134 = P + 11;
	stack[(int)t134] = 10.0;
	
	t135 = P + 12;
	stack[(int)t135] = 10.0;
	
	t138 = P - 0;
	t136 = t138 + 7;
	t137 = stack[(int)t136];
	t141 = P - 0;
	t139 = t141 + 8;
	t140 = stack[(int)t139];
	if(t137 > t140) goto L17;
	goto L18;
	L18:
	t144 = P - 0;
	t142 = t144 + 8;
	t143 = stack[(int)t142];
	t147 = P - 0;
	t145 = t147 + 9;
	t146 = stack[(int)t145];
	if(t143 < t146) goto L19;
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
	t150 = P - 0;
	t148 = t150 + 7;
	t149 = stack[(int)t148];
	t153 = P - 0;
	t151 = t153 + 8;
	t152 = stack[(int)t151];
	if(t149 == t152) goto L22;
	goto L23;
	L22:
	t156 = P - 0;
	t154 = t156 + 11;
	t155 = stack[(int)t154];
	t159 = P - 0;
	t157 = t159 + 12;
	t158 = stack[(int)t157];
	if(t155 == t158) goto L24;
	goto L25;
	L25:
	L23:
	t162 = P - 0;
	t160 = t162 + 9;
	t161 = stack[(int)t160];
	if(14 != t161) goto L26;
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
	t163 = P + 13;
	stack[(int)t163] = 5;
	
	t164 = P + 14;
	stack[(int)t164] = 5;
	
	t165 = P + 15;
	stack[(int)t165] = 100;
	
	t168 = P - 0;
	t166 = t168 + 15;
	t167 = stack[(int)t166];
	t169 = 50 + 50;
	t172 = P - 0;
	t170 = t172 + 13;
	t171 = stack[(int)t170];
	t175 = P - 0;
	t173 = t175 + 13;
	t174 = stack[(int)t173];
	t176 = t171 - t174;
	t177 = t169 + t176;
	if(t167 == t177) goto L29;
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
	t178 = P + 16;
	stack[(int)t178] = 15;
	
	t181 = P - 0;
	t179 = t181 + 16;
	t180 = stack[(int)t179];
	t182 = (int)t180 % (int)2;
	if(t182 == 0) goto L34;
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
	t183 = H;
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
	
	t184 = P + 17;
	t184 = t184 + 1;
	stack[(int)t184] = t183;
	P = P + 17;
	dbrust_printString();
	t185 = stack[(int)P];
	P = P - 17;
	printf("%c", 32);
	
	printf("%c", (char)110);
	printf("%c", (char)105);
	printf("%c", (char)108);
	printf("%c", 32);
	
	printf("%c", 10);

	return 0;
}
