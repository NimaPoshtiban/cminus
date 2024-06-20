#ifndef CFUNCTIONS_H
#define CFUNCTIONS_H
#if defined(_WIN32) || defined(_WIN64)
#include <windows.h>
#include <winuser.h>
#pragma comment (lib, "user32.lib") 
extern void alert(char* msg);
#endif
#endif