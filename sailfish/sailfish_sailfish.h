// +build sailfish sailfish_emulator

#pragma once

#ifndef GO_QTSAILFISH_H
#define GO_QTSAILFISH_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

void* SailfishApp_SailfishApp_Application(int argc, char* argv);
int SailfishApp_SailfishApp_Main(int argc, char* argv);
void* SailfishApp_SailfishApp_CreateView();
void* SailfishApp_SailfishApp_PathTo(char* filename);

#ifdef __cplusplus
}
#endif

#endif