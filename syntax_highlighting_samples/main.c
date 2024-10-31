#include <stdio.h>
#include <string.h>

#include "memory.h"

#define BUFFER_LEN 80

typedef struct {
    int x, y;
} Vec2;

void smallerThan(void) { printf("SMALLER"); }

int main(int argc, char* argv[]) {
    
    Vec2* v;
    v->x = a;
    v->y = a;
    Vec2** c = &v;
    switch (a) {
        case 1:
            break;
    }

    if (!(a >= b) && a <= b || a > b || a == b ||
        a != b) {  // if keyword not highlighted
        smallerThan();
    } else if (a != b) {
        printf("EQUAL\n");  // \n not highlighted
    }

    return 0;
}
