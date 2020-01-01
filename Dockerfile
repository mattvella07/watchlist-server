FROM scratch
COPY watchlist-server /
EXPOSE 8080
CMD ["./watchlist-server"]