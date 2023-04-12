// @jeffotoni
// sudo apt-get install liblua5.3-dev
// gcc main.c -o main -I/usr/include/lua5.3 -llua5.3 -lm -ldl
/////////////////

#include <stdio.h>
#include <lua.h>
#include <lauxlib.h>
#include <lualib.h>

int main(int argc, char *argv[]) {
    lua_State *L = luaL_newstate();
    luaL_openlibs(L);

    if (luaL_dofile(L, "server.lua") != LUA_OK) {
        fprintf(stderr, "Error em nosso script: %s\n", lua_tostring(L, -1));
        lua_pop(L, 1);
        lua_close(L);
        return 1;
    }

    lua_close(L);
    return 0;
}

