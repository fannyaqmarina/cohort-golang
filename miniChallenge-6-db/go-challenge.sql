PGDMP         .            	    {            go-challenge    14.3    14.3     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    25535    go-challenge    DATABASE     q   CREATE DATABASE "go-challenge" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'Indonesian_Indonesia.1252';
    DROP DATABASE "go-challenge";
                postgres    false            �            1259    25556    products    TABLE     �   CREATE TABLE public.products (
    id bigint NOT NULL,
    name character varying(100),
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
    DROP TABLE public.products;
       public         heap    postgres    false            �            1259    25555    products_id_seq    SEQUENCE     x   CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.products_id_seq;
       public          postgres    false    212            �           0    0    products_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;
          public          postgres    false    211            �            1259    25542    variants    TABLE     �   CREATE TABLE public.variants (
    id bigint NOT NULL,
    variant_name character varying,
    quantity integer,
    product_id bigint,
    created_at time with time zone,
    updated_at time with time zone
);
    DROP TABLE public.variants;
       public         heap    postgres    false            �            1259    25541    variants_id_seq    SEQUENCE     x   CREATE SEQUENCE public.variants_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.variants_id_seq;
       public          postgres    false    210            �           0    0    variants_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.variants_id_seq OWNED BY public.variants.id;
          public          postgres    false    209            b           2604    25559    products id    DEFAULT     j   ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);
 :   ALTER TABLE public.products ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    212    211    212            a           2604    25545    variants id    DEFAULT     j   ALTER TABLE ONLY public.variants ALTER COLUMN id SET DEFAULT nextval('public.variants_id_seq'::regclass);
 :   ALTER TABLE public.variants ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    209    210    210            �          0    25556    products 
   TABLE DATA           D   COPY public.products (id, name, created_at, updated_at) FROM stdin;
    public          postgres    false    212   �       �          0    25542    variants 
   TABLE DATA           b   COPY public.variants (id, variant_name, quantity, product_id, created_at, updated_at) FROM stdin;
    public          postgres    false    210   _       �           0    0    products_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.products_id_seq', 5, true);
          public          postgres    false    211                        0    0    variants_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.variants_id_seq', 7, true);
          public          postgres    false    209            f           2606    25561    products products_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.products DROP CONSTRAINT products_pkey;
       public            postgres    false    212            d           2606    25549    variants variants_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.variants
    ADD CONSTRAINT variants_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.variants DROP CONSTRAINT variants_pkey;
       public            postgres    false    210            g           2606    25562    variants product_variant    FK CONSTRAINT     �   ALTER TABLE ONLY public.variants
    ADD CONSTRAINT product_variant FOREIGN KEY (product_id) REFERENCES public.products(id) NOT VALID;
 B   ALTER TABLE ONLY public.variants DROP CONSTRAINT product_variant;
       public          postgres    false    212    3174    210            �   v   x�}�;
�P@���*�C����,�4.�F�S!�i�{Ik��^���6?���DG���$V�	�Y�����,�q"��bG�-��T���ty��{�^�۩��xk�����x�pE � \�+0      �   �   x�m���0Eg�+ؑ"ǉ�abC��A�"�V���I)�!!o����
���2�2��2QI�ʨ=�?HT�n�`>1G$YR��gl��p����S
gM���E[#a��jP�h
�${t��5����K�ol�(Yfi�C��y �o��he�����(�'�̨\���>�     