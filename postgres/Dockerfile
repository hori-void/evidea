FROM postgres:14

# Time Zone
ENV TZ Asia/Tokyo

# Install locales package
RUN apt-get update && apt-get install -y locales && \
    rm -rf /var/lib/apt/lists/*

# Generate ja_JP.UTF-8 locale
RUN sed -i '/ja_JP.UTF-8/s/^# //g' /etc/locale.gen && \
    locale-gen ja_JP.UTF-8

# Set environment variables
ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8

# FROM postgres:14

# # Time ZoneAc
# ENV TZ Asia/Tokyo

# # Language
# ENV LANG ja_JP.UTF-8
# ENV LANGUAGE ja_JP:ja
# ENV LC_ALL ja_JP.UTF-8
# ENV LC_CTYPE=ja_JP.UTF-8
# ENV LC_COLLATE=ja_JP.UTF-8
# ENV LC_MESSAGES=C.UTF-8
# ENV LC_NUMERIC=C.UTF-8
# ENV LC_TIME=C.UTF-8
# ENV LC_MONETARY=C.UTF-8