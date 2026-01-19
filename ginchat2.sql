--
-- PostgreSQL database dump
--

\restrict ibhc0LYhKYR4gVUN342gKDnNjPaycMbyQhiXRny8fUxgREglhUaHFaEyZ6WjcYt

-- Dumped from database version 18.0
-- Dumped by pg_dump version 18.0

-- Started on 2026-01-19 19:06:56

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- TOC entry 226 (class 1259 OID 16544)
-- Name: groups; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.groups (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    creator_id integer NOT NULL,
    created_at timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.groups OWNER TO gorm;

--
-- TOC entry 5050 (class 0 OID 0)
-- Dependencies: 226
-- Name: TABLE groups; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON TABLE public.groups IS '群组表';


--
-- TOC entry 5051 (class 0 OID 0)
-- Dependencies: 226
-- Name: COLUMN groups.id; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.groups.id IS '主键ID';


--
-- TOC entry 5052 (class 0 OID 0)
-- Dependencies: 226
-- Name: COLUMN groups.name; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.groups.name IS '群组名称（唯一）';


--
-- TOC entry 5053 (class 0 OID 0)
-- Dependencies: 226
-- Name: COLUMN groups.creator_id; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.groups.creator_id IS '创建者用户ID';


--
-- TOC entry 5054 (class 0 OID 0)
-- Dependencies: 226
-- Name: COLUMN groups.created_at; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.groups.created_at IS '创建时间';


--
-- TOC entry 5055 (class 0 OID 0)
-- Dependencies: 226
-- Name: COLUMN groups.updated_at; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.groups.updated_at IS '更新时间';


--
-- TOC entry 225 (class 1259 OID 16543)
-- Name: groups_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.groups_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.groups_id_seq OWNER TO gorm;

--
-- TOC entry 5056 (class 0 OID 0)
-- Dependencies: 225
-- Name: groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gorm
--

ALTER SEQUENCE public.groups_id_seq OWNED BY public.groups.id;


--
-- TOC entry 222 (class 1259 OID 16485)
-- Name: message; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.message (
    id bigint NOT NULL,
    created_at timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at bigint,
    from_user_id integer,
    to_user_id integer,
    content character varying(2500) DEFAULT NULL::character varying,
    url character varying(350) DEFAULT NULL::character varying,
    pic text,
    message_type smallint,
    content_type smallint
);


ALTER TABLE public.message OWNER TO gorm;

--
-- TOC entry 5057 (class 0 OID 0)
-- Dependencies: 222
-- Name: TABLE message; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON TABLE public.message IS '消息表';


--
-- TOC entry 5058 (class 0 OID 0)
-- Dependencies: 222
-- Name: COLUMN message.id; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.message.id IS 'id';


--
-- TOC entry 5059 (class 0 OID 0)
-- Dependencies: 222
-- Name: COLUMN message.created_at; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.message.created_at IS '创建时间';


--
-- TOC entry 5060 (class 0 OID 0)
-- Dependencies: 222
-- Name: COLUMN message.updated_at; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.message.updated_at IS '更新时间';


--
-- TOC entry 5061 (class 0 OID 0)
-- Dependencies: 222
-- Name: COLUMN message.deleted_at; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.message.deleted_at IS '删除时间戳';


--
-- TOC entry 5062 (class 0 OID 0)
-- Dependencies: 222
-- Name: COLUMN message.from_user_id; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.message.from_user_id IS '发送人ID';


--
-- TOC entry 5063 (class 0 OID 0)
-- Dependencies: 222
-- Name: COLUMN message.to_user_id; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.message.to_user_id IS '发送对象ID';


--
-- TOC entry 5064 (class 0 OID 0)
-- Dependencies: 222
-- Name: COLUMN message.content; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.message.content IS '消息内容';


--
-- TOC entry 5065 (class 0 OID 0)
-- Dependencies: 222
-- Name: COLUMN message.url; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.message.url IS '文件或者图片地址';


--
-- TOC entry 5066 (class 0 OID 0)
-- Dependencies: 222
-- Name: COLUMN message.pic; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.message.pic IS '缩略图';


--
-- TOC entry 5067 (class 0 OID 0)
-- Dependencies: 222
-- Name: COLUMN message.message_type; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.message.message_type IS '消息类型：1单聊，2群聊';


--
-- TOC entry 5068 (class 0 OID 0)
-- Dependencies: 222
-- Name: COLUMN message.content_type; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.message.content_type IS '消息内容类型：1文字，2语音，3视频';


--
-- TOC entry 221 (class 1259 OID 16484)
-- Name: message_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.message_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.message_id_seq OWNER TO gorm;

--
-- TOC entry 5069 (class 0 OID 0)
-- Dependencies: 221
-- Name: message_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gorm
--

ALTER SEQUENCE public.message_id_seq OWNED BY public.message.id;


--
-- TOC entry 224 (class 1259 OID 16502)
-- Name: persona; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.persona (
    id bigint NOT NULL,
    user_id integer NOT NULL,
    name character varying(100) DEFAULT ''::character varying NOT NULL,
    prompt text NOT NULL,
    created_at timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.persona OWNER TO gorm;

--
-- TOC entry 5070 (class 0 OID 0)
-- Dependencies: 224
-- Name: TABLE persona; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON TABLE public.persona IS '用户角色设定表';


--
-- TOC entry 5071 (class 0 OID 0)
-- Dependencies: 224
-- Name: COLUMN persona.id; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.persona.id IS '主键ID';


--
-- TOC entry 5072 (class 0 OID 0)
-- Dependencies: 224
-- Name: COLUMN persona.user_id; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.persona.user_id IS '用户ID，关联用户表';


--
-- TOC entry 5073 (class 0 OID 0)
-- Dependencies: 224
-- Name: COLUMN persona.name; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.persona.name IS '角色名称';


--
-- TOC entry 5074 (class 0 OID 0)
-- Dependencies: 224
-- Name: COLUMN persona.prompt; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.persona.prompt IS '角色设定提示词';


--
-- TOC entry 5075 (class 0 OID 0)
-- Dependencies: 224
-- Name: COLUMN persona.created_at; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.persona.created_at IS '创建时间';


--
-- TOC entry 5076 (class 0 OID 0)
-- Dependencies: 224
-- Name: COLUMN persona.updated_at; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.persona.updated_at IS '更新时间';


--
-- TOC entry 223 (class 1259 OID 16501)
-- Name: persona_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.persona_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.persona_id_seq OWNER TO gorm;

--
-- TOC entry 5077 (class 0 OID 0)
-- Dependencies: 223
-- Name: persona_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gorm
--

ALTER SEQUENCE public.persona_id_seq OWNED BY public.persona.id;


--
-- TOC entry 228 (class 1259 OID 16565)
-- Name: state; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.state (
    id bigint NOT NULL,
    group_id bigint NOT NULL,
    name character varying(100) NOT NULL,
    creator_id integer NOT NULL,
    prompt text NOT NULL,
    created_at timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.state OWNER TO gorm;

--
-- TOC entry 5078 (class 0 OID 0)
-- Dependencies: 228
-- Name: TABLE state; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON TABLE public.state IS '群组状态定义表';


--
-- TOC entry 5079 (class 0 OID 0)
-- Dependencies: 228
-- Name: COLUMN state.id; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.state.id IS '主键ID';


--
-- TOC entry 5080 (class 0 OID 0)
-- Dependencies: 228
-- Name: COLUMN state.group_id; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.state.group_id IS '所属群组ID';


--
-- TOC entry 5081 (class 0 OID 0)
-- Dependencies: 228
-- Name: COLUMN state.name; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.state.name IS '状态名称（同一群组内唯一）';


--
-- TOC entry 5082 (class 0 OID 0)
-- Dependencies: 228
-- Name: COLUMN state.creator_id; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.state.creator_id IS '创建者用户ID';


--
-- TOC entry 5083 (class 0 OID 0)
-- Dependencies: 228
-- Name: COLUMN state.prompt; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.state.prompt IS '状态作用提示词';


--
-- TOC entry 5084 (class 0 OID 0)
-- Dependencies: 228
-- Name: COLUMN state.created_at; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.state.created_at IS '创建时间';


--
-- TOC entry 5085 (class 0 OID 0)
-- Dependencies: 228
-- Name: COLUMN state.updated_at; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.state.updated_at IS '更新时间';


--
-- TOC entry 227 (class 1259 OID 16564)
-- Name: state_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.state_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.state_id_seq OWNER TO gorm;

--
-- TOC entry 5086 (class 0 OID 0)
-- Dependencies: 227
-- Name: state_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gorm
--

ALTER SEQUENCE public.state_id_seq OWNED BY public.state.id;


--
-- TOC entry 220 (class 1259 OID 16480)
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_id_seq OWNER TO gorm;

--
-- TOC entry 219 (class 1259 OID 16472)
-- Name: user; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public."user" (
    id bigint DEFAULT nextval('public.user_id_seq'::regclass) NOT NULL,
    username character varying(64) NOT NULL,
    password_hash character varying(255) NOT NULL,
    email character varying(128),
    phone character varying(32)
);


ALTER TABLE public."user" OWNER TO gorm;

--
-- TOC entry 230 (class 1259 OID 16595)
-- Name: user_group_state; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.user_group_state (
    id bigint NOT NULL,
    user_id integer NOT NULL,
    group_id bigint NOT NULL,
    state_id bigint NOT NULL,
    assigned_by integer NOT NULL,
    expires_at timestamp(3) without time zone DEFAULT NULL::timestamp without time zone,
    created_at timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.user_group_state OWNER TO gorm;

--
-- TOC entry 5087 (class 0 OID 0)
-- Dependencies: 230
-- Name: TABLE user_group_state; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON TABLE public.user_group_state IS '用户群组状态关联表';


--
-- TOC entry 5088 (class 0 OID 0)
-- Dependencies: 230
-- Name: COLUMN user_group_state.id; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.user_group_state.id IS '主键ID';


--
-- TOC entry 5089 (class 0 OID 0)
-- Dependencies: 230
-- Name: COLUMN user_group_state.user_id; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.user_group_state.user_id IS '用户ID';


--
-- TOC entry 5090 (class 0 OID 0)
-- Dependencies: 230
-- Name: COLUMN user_group_state.group_id; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.user_group_state.group_id IS '群组ID';


--
-- TOC entry 5091 (class 0 OID 0)
-- Dependencies: 230
-- Name: COLUMN user_group_state.state_id; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.user_group_state.state_id IS '状态ID';


--
-- TOC entry 5092 (class 0 OID 0)
-- Dependencies: 230
-- Name: COLUMN user_group_state.assigned_by; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.user_group_state.assigned_by IS '由谁附加该状态';


--
-- TOC entry 5093 (class 0 OID 0)
-- Dependencies: 230
-- Name: COLUMN user_group_state.expires_at; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.user_group_state.expires_at IS '状态过期时间（NULL表示永不过期）';


--
-- TOC entry 5094 (class 0 OID 0)
-- Dependencies: 230
-- Name: COLUMN user_group_state.created_at; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.user_group_state.created_at IS '创建时间';


--
-- TOC entry 5095 (class 0 OID 0)
-- Dependencies: 230
-- Name: COLUMN user_group_state.updated_at; Type: COMMENT; Schema: public; Owner: gorm
--

COMMENT ON COLUMN public.user_group_state.updated_at IS '更新时间';


--
-- TOC entry 229 (class 1259 OID 16594)
-- Name: user_group_state_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.user_group_state_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_group_state_id_seq OWNER TO gorm;

--
-- TOC entry 5096 (class 0 OID 0)
-- Dependencies: 229
-- Name: user_group_state_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gorm
--

ALTER SEQUENCE public.user_group_state_id_seq OWNED BY public.user_group_state.id;


--
-- TOC entry 4844 (class 2604 OID 16547)
-- Name: groups id; Type: DEFAULT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.groups ALTER COLUMN id SET DEFAULT nextval('public.groups_id_seq'::regclass);


--
-- TOC entry 4835 (class 2604 OID 16488)
-- Name: message id; Type: DEFAULT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.message ALTER COLUMN id SET DEFAULT nextval('public.message_id_seq'::regclass);


--
-- TOC entry 4840 (class 2604 OID 16505)
-- Name: persona id; Type: DEFAULT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.persona ALTER COLUMN id SET DEFAULT nextval('public.persona_id_seq'::regclass);


--
-- TOC entry 4847 (class 2604 OID 16568)
-- Name: state id; Type: DEFAULT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.state ALTER COLUMN id SET DEFAULT nextval('public.state_id_seq'::regclass);


--
-- TOC entry 4850 (class 2604 OID 16598)
-- Name: user_group_state id; Type: DEFAULT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.user_group_state ALTER COLUMN id SET DEFAULT nextval('public.user_group_state_id_seq'::regclass);


--
-- TOC entry 4871 (class 2606 OID 16554)
-- Name: groups groups_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (id);


--
-- TOC entry 4862 (class 2606 OID 16497)
-- Name: message message_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.message
    ADD CONSTRAINT message_pkey PRIMARY KEY (id);


--
-- TOC entry 4867 (class 2606 OID 16516)
-- Name: persona persona_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.persona
    ADD CONSTRAINT persona_pkey PRIMARY KEY (id);


--
-- TOC entry 4879 (class 2606 OID 16579)
-- Name: state state_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.state
    ADD CONSTRAINT state_pkey PRIMARY KEY (id);


--
-- TOC entry 4875 (class 2606 OID 16556)
-- Name: groups uk_groups_name; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT uk_groups_name UNIQUE (name);


--
-- TOC entry 4869 (class 2606 OID 16521)
-- Name: persona uk_persona_user_name; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.persona
    ADD CONSTRAINT uk_persona_user_name UNIQUE (user_id, name);


--
-- TOC entry 4881 (class 2606 OID 16581)
-- Name: state uk_state_group_name; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.state
    ADD CONSTRAINT uk_state_group_name UNIQUE (group_id, name);


--
-- TOC entry 4888 (class 2606 OID 16610)
-- Name: user_group_state uk_user_group_state; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.user_group_state
    ADD CONSTRAINT uk_user_group_state UNIQUE (user_id, group_id, state_id);


--
-- TOC entry 4890 (class 2606 OID 16608)
-- Name: user_group_state user_group_state_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.user_group_state
    ADD CONSTRAINT user_group_state_pkey PRIMARY KEY (id);


--
-- TOC entry 4855 (class 2606 OID 16479)
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- TOC entry 4857 (class 2606 OID 16483)
-- Name: user user_username_unique; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_username_unique UNIQUE (username);


--
-- TOC entry 4872 (class 1259 OID 16563)
-- Name: idx_groups_created_at; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_groups_created_at ON public.groups USING btree (created_at);


--
-- TOC entry 4873 (class 1259 OID 16562)
-- Name: idx_groups_creator_id; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_groups_creator_id ON public.groups USING btree (creator_id);


--
-- TOC entry 4858 (class 1259 OID 16498)
-- Name: idx_message_deleted_at; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_message_deleted_at ON public.message USING btree (deleted_at);


--
-- TOC entry 4859 (class 1259 OID 16499)
-- Name: idx_message_from_user_id; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_message_from_user_id ON public.message USING btree (from_user_id);


--
-- TOC entry 4860 (class 1259 OID 16500)
-- Name: idx_message_to_user_id; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_message_to_user_id ON public.message USING btree (to_user_id);


--
-- TOC entry 4863 (class 1259 OID 16518)
-- Name: idx_persona_created_at; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_persona_created_at ON public.persona USING btree (created_at);


--
-- TOC entry 4864 (class 1259 OID 16519)
-- Name: idx_persona_name; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_persona_name ON public.persona USING btree (name);


--
-- TOC entry 4865 (class 1259 OID 16517)
-- Name: idx_persona_user_id; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_persona_user_id ON public.persona USING btree (user_id);


--
-- TOC entry 4876 (class 1259 OID 16593)
-- Name: idx_state_creator_id; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_state_creator_id ON public.state USING btree (creator_id);


--
-- TOC entry 4877 (class 1259 OID 16592)
-- Name: idx_state_group_id; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_state_group_id ON public.state USING btree (group_id);


--
-- TOC entry 4882 (class 1259 OID 16633)
-- Name: idx_user_group_state_assigner; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_user_group_state_assigner ON public.user_group_state USING btree (assigned_by);


--
-- TOC entry 4883 (class 1259 OID 16635)
-- Name: idx_user_group_state_created_at; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_user_group_state_created_at ON public.user_group_state USING btree (created_at);


--
-- TOC entry 4884 (class 1259 OID 16634)
-- Name: idx_user_group_state_expires; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_user_group_state_expires ON public.user_group_state USING btree (expires_at);


--
-- TOC entry 4885 (class 1259 OID 16632)
-- Name: idx_user_group_state_state; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_user_group_state_state ON public.user_group_state USING btree (state_id);


--
-- TOC entry 4886 (class 1259 OID 16631)
-- Name: idx_user_group_state_user_group; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_user_group_state_user_group ON public.user_group_state USING btree (user_id, group_id);


--
-- TOC entry 4891 (class 2606 OID 16557)
-- Name: groups fk_groups_creator; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT fk_groups_creator FOREIGN KEY (creator_id) REFERENCES public."user"(id) ON DELETE RESTRICT;


--
-- TOC entry 4892 (class 2606 OID 16587)
-- Name: state fk_state_creator; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.state
    ADD CONSTRAINT fk_state_creator FOREIGN KEY (creator_id) REFERENCES public."user"(id) ON DELETE RESTRICT;


--
-- TOC entry 4893 (class 2606 OID 16582)
-- Name: state fk_state_group; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.state
    ADD CONSTRAINT fk_state_group FOREIGN KEY (group_id) REFERENCES public.groups(id) ON DELETE CASCADE;


--
-- TOC entry 4894 (class 2606 OID 16626)
-- Name: user_group_state fk_user_group_state_assigner; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.user_group_state
    ADD CONSTRAINT fk_user_group_state_assigner FOREIGN KEY (assigned_by) REFERENCES public."user"(id) ON DELETE RESTRICT;


--
-- TOC entry 4895 (class 2606 OID 16616)
-- Name: user_group_state fk_user_group_state_group; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.user_group_state
    ADD CONSTRAINT fk_user_group_state_group FOREIGN KEY (group_id) REFERENCES public.groups(id) ON DELETE CASCADE;


--
-- TOC entry 4896 (class 2606 OID 16621)
-- Name: user_group_state fk_user_group_state_state; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.user_group_state
    ADD CONSTRAINT fk_user_group_state_state FOREIGN KEY (state_id) REFERENCES public.state(id) ON DELETE CASCADE;


--
-- TOC entry 4897 (class 2606 OID 16611)
-- Name: user_group_state fk_user_group_state_user; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.user_group_state
    ADD CONSTRAINT fk_user_group_state_user FOREIGN KEY (user_id) REFERENCES public."user"(id) ON DELETE CASCADE;


-- Completed on 2026-01-19 19:06:56

--
-- PostgreSQL database dump complete
--

\unrestrict ibhc0LYhKYR4gVUN342gKDnNjPaycMbyQhiXRny8fUxgREglhUaHFaEyZ6WjcYt

