#include <stdio.h>
#include "libsunny.h"
int main(){
    GoString str;
	GoUint8 show;
    str = Hello();   
    Test();
    printf("%lld\n",str.n);
	show=IsOrientationZero(str);
	printf("%d\n",show);
	return 0;
}

//gcc main.c -o test -I./ -L./ -lsunny