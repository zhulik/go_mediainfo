#include <wchar.h>
#include <MediaInfoDLL/MediaInfoDLL.h>
#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <locale.h>

const wchar_t *toWchar(const char *c)
{
    const size_t cSize = strlen(c)+1;
    wchar_t* wc = malloc(cSize * sizeof(wchar_t));
    mbstowcs (wc, c, cSize);
    return wc;
}

const char *toChar(const wchar_t *c)
{
    const size_t cSize = wcslen(c)+1;
    char* wc = malloc(cSize * sizeof(char));
    wcstombs(wc, c, cSize);
    return wc;
}

void *GoMediaInfo_New() {
    setlocale(LC_CTYPE, "");
    return MediaInfo_New();
}

void GoMediaInfo_Delete(void *handle) {
    MediaInfo_Delete(handle);
}

size_t GoMediaInfo_Open(void *handle, char *name) {
    return MediaInfo_Open(handle, toWchar(name));
}

void GoMediaInfo_Close(void *handle) {
    MediaInfo_Close(handle);
}

const char *GoMediaInfoGet(void *handle, char *name) {
    return toChar(MediaInfo_Get(handle, MediaInfo_Stream_General, 0,  toWchar(name), MediaInfo_Info_Text, MediaInfo_Info_Name));
}