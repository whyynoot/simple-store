BEGIN;

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = ON;
SET check_function_bodies = FALSE;
SET client_min_messages = WARNING;
SET search_path = public, extensions;
SET default_tablespace = '';
SET default_with_oids = FALSE;

-- EXTENSIONS --

CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- TABLES --


CREATE TABLE public.category
(
    id   SERIAL PRIMARY KEY,
    name_ru TEXT NOT NULL,
    api_name TEXT NOT NULL UNIQUE
);

CREATE TABLE public.product
(
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name_ru       TEXT NOT NULL,
    api_name      TEXT NOT NULL UNIQUE,
    description   TEXT NOT NULL,
    image_path    TEXT NOT NULL,
    price         BIGINT NOT NULL,
    category_id   SERIAL references public.category,
    CONSTRAINT positive_price CHECK (price > 0)
);

CREATE TABLE public.order
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    phone       TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO public.category (name_ru, api_name) VALUES ('Волосы', 'hair');
INSERT INTO public.category (name_ru, api_name) VALUES ('Тело', 'body');

-- DATA --
INSERT INTO public.product (id, name_ru, api_name, description, image_path, price, category_id)
VALUES (
           DEFAULT,
           'Шампунь',
           'shampo',
           'Для окрашенных волос',
           'shampo.jpg',
           400,
            1
       );

INSERT INTO public.product (id, name_ru, api_name, description, image_path, price, category_id)
VALUES (
           DEFAULT,
           'Молочко для тела',
           'milk',
           'Увлажняет кожу со вкусом кокоса',
           'milk.jpg',
           330,
           2
       );

INSERT INTO public.product (id, name_ru, api_name, description, image_path, price, category_id)
VALUES (
           DEFAULT,
           'Кондиционер для волос',
           'condo',
           'Предаст волосам мягкость и сохранит яркость цвета',
           'condo.jpg',
           330,
           1
       );

COMMIT;