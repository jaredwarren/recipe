FROM       scratch
ENV        PORT 80
EXPOSE     80
ADD        recipe recipe
ENTRYPOINT ["/recipe"]