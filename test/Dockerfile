FROM local/test/env-mapper as mapper
FROM asecurityteam/sdcli:v1
USER root
RUN pip3 install pytest-testinfra==7.0.0
RUN mkdir /test
ADD *.py /test/
COPY --from=mapper /app /env-mapper
WORKDIR /test/
ENTRYPOINT ["py.test", "-v"]