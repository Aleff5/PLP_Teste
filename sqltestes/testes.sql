CREATE TABLE Herois (
    id_heroi SERIAL PRIMARY KEY,
    nome_heroi VARCHAR(50) NOT NULL,
    nome_real VARCHAR(50) NOT NULL,
    sexo VARCHAR(10) NOT NULL,
    altura FLOAT,
    local_nascimento VARCHAR(100) NOT NULL,
    data_nascimento DATE,
    peso FLOAT,
    popularidade INT NOT NULL CHECK (popularidade BETWEEN 0 AND 100),
    forca INT NOT NULL CHECK (forca BETWEEN 0 AND 100),
    status_atividade VARCHAR(20) CHECK (status_atividade IN ('Ativo', 'Banido', 'Inativo'))
    esconder BOOLEAN NOT NULL DEFAULT FALSE
);


CREATE TABLE Poderes (
    id_poder SERIAL NOT NULL PRIMARY KEY,
    poder VARCHAR(50),
    descricao VARCHAR(255)
);



CREATE TABLE Batalhas (
    id_batalha SERIAL PRIMARY KEY,
    local VARCHAR(100) NOT NULL,
    data DATE NOT NULL,
    descricao VARCHAR(255),
    resultado VARCHAR(20) CHECK (resultado IN ('Sucesso', 'Fracasso'))
);

CREATE TABLE Herois_Batalhas (
    id_heroi INT NOT NULL,
    id_batalha INT NOT NULL,
    CONSTRAINT pk_heroi_batalha PRIMARY KEY (id_heroi, id_batalha),
    CONSTRAINT fk_heroi_batalha_heroi FOREIGN KEY (id_heroi) REFERENCES Herois(id_heroi),
    CONSTRAINT fk_heroi_batalha_batalha FOREIGN KEY (id_batalha) REFERENCES Batalhas(id_batalha)
);

CREATE TABLE Missoes (
    id_missao SERIAL PRIMARY KEY,
    nome_missao VARCHAR(100) NOT NULL,
    descricao VARCHAR(255),
    nivel_dificuldade INT NOT NULL CHECK (nivel_dificuldade BETWEEN 1 AND 10),
    resultado VARCHAR(20) CHECK (resultado IN ('Sucesso', 'Fracasso')),
    recompensa VARCHAR(50) -- ex: "+10 popularidade", "+5 força"   
);

CREATE TABLE Herois_Missoes (
    id_heroi INT NOT NULL,
    id_missao INT NOT NULL,
    CONSTRAINT pk_heroi_missao PRIMARY KEY (id_heroi, id_missao),
    CONSTRAINT fk_heroi_missao_heroi FOREIGN KEY (id_heroi) REFERENCES Herois(id_heroi),
    CONSTRAINT fk_heroi_missao_missao FOREIGN KEY (id_missao) REFERENCES Missoes(id_missao)
);


CREATE TABLE Crimes (
    id_crime SERIAL PRIMARY KEY,
    nome_crime VARCHAR(100) NOT NULL,
    severidade INT NOT NULL CHECK (severidade BETWEEN 1 AND 10)
);

--alteração: criação de id_ocorrecia como chave primaria.
CREATE TABLE Herois_Crimes (
    id_ocorrencia SERIAL PRIMARY KEY,
    id_heroi INT NOT NULL,
    id_crime INT NOT NULL,
    data_crime DATE NOT NULL,
    descricao_evento VARCHAR(255),
    esconder BOOLEAN NOT NULL DEFAULT FALSE,
    CONSTRAINT fk_heroi_crime_heroi FOREIGN KEY (id_heroi) REFERENCES Herois(id_heroi),
    CONSTRAINT fk_heroi_crime_crime FOREIGN KEY (id_crime) REFERENCES Crimes(id_crime)
);

-- trigger para atualizar status heroi baseado na popularidade
CREATE OR REPLACE FUNCTION att_heroi_status_func()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.popularidade < 20 THEN
        NEW.status_atividade := 'Banido';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER att_heroi_status
BEFORE UPDATE ON Herois
FOR EACH ROW
EXECUTE FUNCTION att_heroi_status_func();







-- trigger para atualizar popularidade do heroi baseado nos crimes cometidos
CREATE OR REPLACE FUNCTION att_popularidade_heroi_func()
RETURNS TRIGGER AS $$
DECLARE
    severidade INT;
    reducao FLOAT;
BEGIN
    -- Obter a severidade do crime registrado
    SELECT severidade INTO severidade FROM Crimes WHERE id_crime = NEW.id_crime;

    -- Determinar a redução na popularidade com base na severidade
    IF severidade BETWEEN 1 AND 4 THEN
        reducao := 0.15;
    ELSIF severidade BETWEEN 5 AND 8 THEN
        reducao := 0.20;
    ELSIF severidade BETWEEN 9 AND 10 THEN
        reducao := 0.50;
    END IF;

    -- Atualizar a popularidade do herói aplicando a redução calculada
    UPDATE Herois
    SET popularidade = GREATEST(0, popularidade * (1 - reducao))
    WHERE id_heroi = NEW.id_heroi;

    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER att_popularidade_heroi
AFTER INSERT ON Herois_Crimes
FOR EACH ROW
EXECUTE FUNCTION att_popularidade_heroi_func();





-- trigger para ajustar atributos apos missao
CREATE OR REPLACE FUNCTION ajustar_atributos_missao_func()
RETURNS TRIGGER AS $$
DECLARE
    resultado VARCHAR(20);
    aumento_forca FLOAT;
    aumento_popularidade FLOAT;
BEGIN
    -- Obter o resultado da missão
    SELECT m.resultado INTO resultado FROM Missoes m WHERE m.id_missao = NEW.id_missao;

    -- Se a missão for um sucesso
    IF resultado = 'Sucesso' THEN
        aumento_forca := 0.10;          -- 10% de aumento na força
        aumento_popularidade := 0.15;   -- 15% de aumento na popularidade

        -- Aumenta a força e a popularidade do herói
        UPDATE Herois
        SET 
            forca = LEAST(100, forca * (1 + aumento_forca)),  -- Limite máximo de 100 para a força
            popularidade = LEAST(100, popularidade * (1 + aumento_popularidade)) -- Limite máximo de 100 para a popularidade
        WHERE id_heroi = NEW.id_heroi;

    -- Se a missão for um fracasso
    ELSIF resultado = 'Fracasso' THEN
        aumento_popularidade := 0.10;    -- 10% de redução na popularidade

        -- Reduz a popularidade do herói
        UPDATE Herois
        SET 
            popularidade = GREATEST(0, popularidade * (1 - aumento_popularidade)) -- Limite mínimo de 0
        WHERE id_heroi = NEW.id_heroi;
    END IF;

    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER ajustar_atributos_missao
AFTER INSERT ON Herois_Missoes
FOR EACH ROW
EXECUTE FUNCTION ajustar_atributos_missao_func();


CREATE TABLE Herois_Poderes(
    id_heroi INT NOT NULL,
    id_poder INT NOT NULL,
    CONSTRAINT pk_heroi_poder PRIMARY KEY (id_heroi, id_poder),
    CONSTRAINT fk_heroi_poder FOREIGN KEY (id_heroi) REFERENCES Herois(id_heroi),
    CONSTRAINT fk_poder_heroi FOREIGN KEY (id_poder) REFERENCES Poder(id_poder)
)