# WIP Toplevel makefile
all:
	(mkdir -p build/assets)
	(cd web; yarn exec vuepress build docs; mv ./docs/.vuepress/dist/* ../build/assets; cd ..)
	(cd api; make clean; cp -R ../build/assets .;  make generate build; mv bin/api ../build; cd ..)

clean:
	rm -rf ./build
