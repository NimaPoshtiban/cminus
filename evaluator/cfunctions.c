#include "cfunctions.h"



extern void alert(char* msg){
    MessageBoxA(0, msg, "Success",MB_ICONINFORMATION| MB_OK);
}