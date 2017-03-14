FROM scratch

ADD reaper /reaper

ENTRYPOINT ["/reaper"]
