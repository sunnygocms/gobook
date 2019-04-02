#include <stdio.h>
#include "libsunny.h"
int main(){
//    GoString str;
	GoUint8 show;
//    str = Hello();   
//    Test();
//    printf("%lld\n",str.n);
	GoString name = {"1.jpg", 5};
	show=IsOrientationZero(name);
	printf("%d\n",show);
	// 横着的
	GoString name2 = {"2.jpg", 5};
	show=IsOrientationZero(name2);
	printf("%d\n",show);
	return 0;
}

//gcc main.c -o test -I./ -L./ -lsunny