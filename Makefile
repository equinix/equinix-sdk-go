*:
  service_name=$@
  mkdir -p services/$@
  cp Makefile.sdk.template services/$@/Makefile
  cd $@ && make
