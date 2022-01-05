# /bin/sh

if [ ! -d /docs ]; then
  docsify init ./docs
  docsify serve docs
fi

docsify init ./tmp-docs

if [ ! -f /docs/index.html ]; then
  mv /tmp-docs/index.html /docs
fi

if [ ! -f /docs/README.md ]; then
  mv /tmp-docs/README.md /docs
fi

if [ ! -f /docs/.nojekyll ]; then
  mv /tmp-docs/.nojekyll /docs
fi

docsify serve docs
