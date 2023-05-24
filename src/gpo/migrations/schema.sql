--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3
-- Dumped by pg_dump version 14.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: companies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.companies (
    id uuid NOT NULL,
    city character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    postalnr character varying(255) NOT NULL,
    street character varying(255) NOT NULL,
    streetnr character varying(255) NOT NULL,
    gln character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.companies OWNER TO postgres;

--
-- Name: companies_gpoes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.companies_gpoes (
    id uuid NOT NULL,
    company_id uuid NOT NULL,
    gpo_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.companies_gpoes OWNER TO postgres;

--
-- Name: contractors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contractors (
    id uuid NOT NULL,
    provider character varying(255) NOT NULL,
    provider_id character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.contractors OWNER TO postgres;

--
-- Name: contractors_companies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contractors_companies (
    id uuid NOT NULL,
    contractor_id uuid NOT NULL,
    company_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.contractors_companies OWNER TO postgres;

--
-- Name: contractors_gpoes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contractors_gpoes (
    id uuid NOT NULL,
    contractor_id uuid NOT NULL,
    gpo_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.contractors_gpoes OWNER TO postgres;

--
-- Name: gpoes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.gpoes (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.gpoes OWNER TO postgres;

--
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    is_gdsn boolean NOT NULL,
    gdsn_country_code character varying(255),
    gdsn_gtin character varying(255),
    gdsn_gln character varying(255),
    supplier_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.products OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: suppliercertificates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.suppliercertificates (
    id uuid NOT NULL,
    issuer character varying(255) NOT NULL,
    issued_at timestamp without time zone NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    data_link character varying(255) NOT NULL,
    supplier_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.suppliercertificates OWNER TO postgres;

--
-- Name: supplierreviews; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.supplierreviews (
    id uuid NOT NULL,
    metric character varying(255) NOT NULL,
    min_value integer NOT NULL,
    max_value integer NOT NULL,
    value integer NOT NULL,
    user_id uuid NOT NULL,
    supplier_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.supplierreviews OWNER TO postgres;

--
-- Name: suppliers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.suppliers (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    gln text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.suppliers OWNER TO postgres;

--
-- Name: suppliers_gpoes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.suppliers_gpoes (
    id uuid NOT NULL,
    supplier_id uuid NOT NULL,
    gpo_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.suppliers_gpoes OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    provider character varying(255) NOT NULL,
    provider_id character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    company_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: companies_gpoes companies_gpoes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.companies_gpoes
    ADD CONSTRAINT companies_gpoes_pkey PRIMARY KEY (company_id, gpo_id);


--
-- Name: companies companies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.companies
    ADD CONSTRAINT companies_pkey PRIMARY KEY (id);


--
-- Name: contractors_companies contractors_companies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contractors_companies
    ADD CONSTRAINT contractors_companies_pkey PRIMARY KEY (contractor_id, company_id);


--
-- Name: contractors_gpoes contractors_gpoes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contractors_gpoes
    ADD CONSTRAINT contractors_gpoes_pkey PRIMARY KEY (contractor_id, gpo_id);


--
-- Name: contractors contractors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contractors
    ADD CONSTRAINT contractors_pkey PRIMARY KEY (id);


--
-- Name: gpoes gpoes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gpoes
    ADD CONSTRAINT gpoes_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: suppliercertificates suppliercertificates_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.suppliercertificates
    ADD CONSTRAINT suppliercertificates_pkey PRIMARY KEY (id);


--
-- Name: supplierreviews supplierreviews_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.supplierreviews
    ADD CONSTRAINT supplierreviews_pkey PRIMARY KEY (id);


--
-- Name: suppliers_gpoes suppliers_gpoes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.suppliers_gpoes
    ADD CONSTRAINT suppliers_gpoes_pkey PRIMARY KEY (supplier_id, gpo_id);


--
-- Name: suppliers suppliers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.suppliers
    ADD CONSTRAINT suppliers_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: products products_supplier_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_supplier_id_fkey FOREIGN KEY (supplier_id) REFERENCES public.suppliers(id);


--
-- Name: suppliercertificates suppliercertificates_supplier_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.suppliercertificates
    ADD CONSTRAINT suppliercertificates_supplier_id_fkey FOREIGN KEY (supplier_id) REFERENCES public.suppliers(id);


--
-- Name: supplierreviews supplierreviews_supplier_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.supplierreviews
    ADD CONSTRAINT supplierreviews_supplier_id_fkey FOREIGN KEY (supplier_id) REFERENCES public.suppliers(id);


--
-- Name: supplierreviews supplierreviews_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.supplierreviews
    ADD CONSTRAINT supplierreviews_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: users users_company_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_company_id_fkey FOREIGN KEY (company_id) REFERENCES public.companies(id);


--
-- PostgreSQL database dump complete
--

