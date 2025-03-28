FROM postgres:16

# Configurar las variables de entorno
ENV POSTGRES_USER=usuario
ENV POSTGRES_PASSWORD=secret
ENV POSTGRES_DB=series_tracker

# Copiar el script SQL de inicialización al contenedor
COPY ./init.sql /docker-entrypoint-initdb.d/init.sql

# Exponer el puerto de PostgreSQL
EXPOSE 5432

# Comando de inicio en modo foreground
ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["postgres"]
