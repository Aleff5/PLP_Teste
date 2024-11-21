-- Inserir HEROIS no banco de dados

INSERT INTO Herois (nome_heroi, nome_real, sexo, altura, local_nasc, data_nasc, peso, popularidade, forca, status_atividade)
VALUES
('Homelander', 'John', 'Masculino', 1.91, 'Estados Unidos', '1982-06-10', 90.0, 95, 100, 'Ativo'),
('Starlight', 'Annie January', 'Feminino', 1.65, 'Estados Unidos', '1991-05-01', 55.0, 85, 75, 'Ativo'),
('Queen Maeve', 'Maeve', 'Feminino', 1.75, 'Estados Unidos', '1980-04-15', 70.0, 90, 90, 'Inativo'),
('A-Train', 'Reggie Franklin', 'Masculino', 1.77, 'Estados Unidos', '1986-03-01', 80.0, 70, 65, 'Banido'),
('The Deep', 'Kevin Moskowitz', 'Masculino', 1.80, 'Estados Unidos', '1986-07-25', 85.0, 60, 55, 'Ativo'),
('Black Noir', 'Desconhecido', 'Masculino', 1.88, 'Desconhecido', NULL, 100.0, 80, 95, 'Inativo'),
('The Soldier Boy', 'Ben', 'Masculino', 1.85, 'Estados Unidos', '1940-12-01', 95.0, 75, 90, 'Banido'),
('Kimiko', 'Kimiko', 'Feminino', 1.65, 'Japão', '1985-08-10', 50.0, 50, 85, 'Ativo'),
('Mothers Milk', 'Marvin T. Milk', 'Masculino', 1.80, 'Estados Unidos', '1983-01-15', 90.0, 65, 70, 'Ativo'),
('Frenchie', 'Serge', 'Masculino', 1.75, 'França', '1980-06-20', 75.0, 60, 60, 'Ativo');


-- Inserir PODERES no banco de dados

INSERT INTO Poderes (poder, descricao) VALUES
    ('Super força', 'Força física sobre-humana, capaz de levantar objetos extremamente pesados.'),
    ('Invulnerabilidade', 'Resistência a danos físicos, incluindo balas, explosões e impactos.'),
    ('Visão de calor', 'Emissão de feixes de calor concentrado pelos olhos.'),
    ('Voo', 'Capacidade de voar e se mover pelo ar em alta velocidade.'),
    ('Velocidade sobre-humana', 'Habilidade de se mover a velocidades incríveis, muito acima da média humana.'),
    ('Agilidade aprimorada', 'Reflexos e movimentos rápidos e precisos, superiores a qualquer atleta.'),
    ('Regeneração', 'Capacidade de curar feridas e regenerar partes do corpo rapidamente.'),
    ('Controle mental', 'Habilidade de manipular os pensamentos e ações de outras pessoas.'),
    ('Controle de eletricidade', 'Capacidade de gerar, manipular e lançar descargas elétricas.'),
    ('Transparência', 'Capacidade de ficar invisível, tornando-se indetectável a olho nu.'),
    ('Manipulação de água', 'Habilidade de controlar e moldar a água para criar formas ou ataques.'),
    ('Transformação corporal', 'Capacidade de mudar a forma física ou aparência do corpo.'),
    ('Audição ampliada', 'Capacidade de ouvir sons a grandes distâncias e frequências imperceptíveis aos humanos.'),
    ('Controle de animais', 'Habilidade de se comunicar e controlar animais.'),
    ('Visão de raio-X', 'Capacidade de enxergar através de objetos sólidos.'),
    ('Criação de explosões', 'Habilidade de gerar explosões de energia com as mãos ou corpo.'),
    ('Controle de fogo', 'Habilidade de criar e manipular chamas.'),
    ('Teletransporte', 'Capacidade de se mover instantaneamente de um lugar para outro.'),
    ('Manipulação de energia', 'Habilidade de controlar e redirecionar energia em diferentes formas.'),
    ('Resistência extrema', 'Capacidade de suportar longos períodos de atividade física sem se cansar.'),
    ('Controle de gravidade', 'Habilidade de alterar a gravidade em uma área específica.'),
    ('Telecinese', 'Capacidade de mover objetos com a mente.'),
    ('Criocinese', 'Habilidade de gerar e manipular gelo ou temperaturas extremamente frias.'),
    ('Manipulação de plantas', 'Habilidade de controlar e fazer plantas crescerem rapidamente.'),
    ('Camuflagem', 'Capacidade de se misturar com o ambiente, tornando-se quase invisível.'),
    ('Sentidos ampliados', 'Audição, visão, olfato e tato extremamente aguçados.'),
    ('Paralisia', 'Capacidade de paralisar inimigos com um toque ou olhar.'),
    ('Habilidade de cura', 'Capacidade de curar outras pessoas ao tocá-las.'),
    ('Absorção de energia', 'Habilidade de absorver energia de ataques ou fontes externas para ficar mais forte.');

-- Inserir a tabela Herois_Poderes no banco de dados

INSERT INTO herois_poderes (id_heroi, id_poder) VALUES
(1, 13), -- Homelander - Super força
(1, 14), -- Homelander - Invulnerabilidade
(1, 15), -- Homelander - Visão de calor
(1, 16), -- Homelander - Voo
(2, 14), -- Starlight - Invulnerabilidade
(2, 26), -- Starlight - Controle de eletricidade
(2, 20), -- Starlight - Resistência extrema
(3, 13), -- Queen Maeve - Super força
(3, 14), -- Queen Maeve - Invulnerabilidade
(3, 20), -- Queen Maeve - Resistência extrema
(4, 17), -- A-Train - Velocidade sobre-humana
(4, 17), -- A-Train - Agilidade aprimorada
(5, 27), -- The Deep - Controle de animais
(5, 11), -- The Deep - Manipulação de água
(6, 14), -- Black Noir - Invulnerabilidade
(6, 10), -- Black Noir - Transparência
(7, 13), -- Soldier Boy - Super força
(7, 20), -- Soldier Boy - Resistência extrema
(7, 41), -- Soldier Boy - Absorção de energia
(8, 7),  -- Kimiko - Regeneração
(8, 13), -- Kimiko - Super força
(8, 17), -- Kimiko - Agilidade aprimorada
(9, 13), -- Mothers Milk - Super força
(9, 20), -- Mothers Milk - Resistência extrema
(10, 17), -- Frenchie - Agilidade aprimorada
(10, 26); -- Frenchie - Sentidos ampliados



-- Inserir CRIMES no banco de dados

INSERT INTO Crimes (nome_crime, severidade)
VALUES
('Assassinato', 10),
('Roubo', 7),
('Fraude', 6),
('Sequestro', 9),
('Corrupção', 8),
('Tráfico de drogas', 9),
('Agressão física', 6),
('Estupro', 10),
('Vandalismo', 5),
('Extorsão', 8),
('Hackerismo', 7),
('Assédio sexual', 8),
('Terrorismo', 10),
('Tráfico de seres humanos', 10),
('Falsificação de documentos', 6),
('Lavagem de dinheiro', 9),
('Espionagem', 8),
('Concussão', 5),
('Fraude fiscal', 7),
('Roubos à mão armada', 9),
('Urinar em local público', 3),
('Pedofilia', 8),
('11-09', 10),
('Homofobia', 1),
('Racismo', 7),
('Trafico de animais', 5);

--Inserir as missoes

INSERT INTO Missoes (nome_missao, descricao, nivel_dificuldade)
VALUES
('Caçada ao Supers', 'Investigar e capturar um super que age fora da lei', 7),
('Missão no Submundo', 'Infiltrar uma rede criminosa ligada a supers', 8),
('Sabotagem Corporativa', 'Descobrir e expor segredos da Vought', 6),
('Resgate em Perigo', 'Salvar civis de um ataque descontrolado', 5),
('Operação Nocturna', 'Vigiar um super suspeito durante a noite', 3),
('Confronto Público', 'Confrontar um super em um evento televisivo', 9),
('Negociações Perigosas', 'Medir forças diplomáticas com a Vought', 7),
('Neutralizar Supers', 'Desarmar um super sem causar mortes', 8),
('Proteger Testemunha', 'Escoltar um ex-super disposto a testemunhar', 4),
('Hackeamento Crítico', 'Roubar informações sigilosas da Vought', 6),
('Reconhecimento Urbano', 'Mapear atividades suspeitas de supers na cidade', 2),
('Monitoramento Secreto', 'Acompanhar um super sem ser notado', 3),
('Entrega Segura', 'Transportar um pacote crítico sem chamar atenção', 4),
('Contato Inicial', 'Estabelecer comunicação com uma possível testemunha', 1),
('Investigação Local', 'Coletar evidências em uma cena suspeita', 3),
('Operação Silenciosa', 'Entrar e sair de um prédio sem ser detectado', 5),
('Reunião Clandestina', 'Participar de um encontro secreto sem atrair suspeitas', 4),
('Resgate Rápido', 'Retirar civis de uma área de risco antes da chegada de um super', 2),
('Alerta Comunitário', 'Informar moradores sobre a presença de um super perigoso', 1),
('Teste de Equipamento', 'Avaliar novos dispositivos contra supers', 3);

--inserir na tabela Heroi_Crime

INSERT INTO Heroi_Crime (id_heroi, id_crime, data_crime, descricao_evento, esconder)
VALUES
-- Homelander
(1, 10, '2023-05-12', 'Ataque terrorista em um prédio civil', TRUE),
(1, 7, '2022-08-03', 'Desvio de verba da Vought', TRUE),
(1, 6, '2023-01-15', 'Sequestro de filho de Becca Butcher', TRUE),
(1, 9, '2022-07-22', 'Envolvimento com o tráfico de drogas', TRUE),
(1, 8, '2023-04-18', 'Agressão sexual contra uma funcionária da Vought', TRUE),

-- Starlight
(2, 7, '2023-03-10', 'Exposição de informações sigilosas da Vought', FALSE),
(2, 5, '2022-11-20', 'Destruição de propriedade durante uma luta', FALSE),

-- Queen Maeve
(3, 10, '2020-06-01', 'Participação na queda de um avião', TRUE),
(3, 6, '2021-12-10', 'Sequestro de civis para controle da narrativa da Vought', TRUE),
(3, 9, '2021-03-15', 'Agressão física a criminosos em um bar', FALSE),

-- A-Train
(4, 7, '2023-06-30', 'Uso de Composto V para doping', TRUE),
(4, 9, '2022-01-05', 'Tráfico de drogas envolvendo Composto V', TRUE),
(4, 6, '2023-02-14', 'Agressão a civis após discussão', FALSE),

-- The Deep
(5, 8, '2021-07-12', 'Assédio sexual a funcionária', TRUE),
(5, 10, '2022-03-25', 'Colaboração com atividades terroristas contra ecossistemas', TRUE),
(5, 5, '2022-10-09', 'Vandalismo em protesto ambiental', FALSE),

-- Black Noir
(6, 10, '2019-04-10', 'Atentado terrorista a mando da Vought', TRUE),
(6, 6, '2020-02-20', 'Sequestro de cientista para experimentos', TRUE),

-- The Soldier Boy
(7, 10, '1984-09-14', 'Atrocidades em missões militares secretas', TRUE),
(7, 7, '1986-11-01', 'Fraude em manipulação de propaganda de guerra', TRUE),

-- Kimiko
(8, 5, '2023-08-15', 'Dano a propriedades durante combate', FALSE),
(8, 9, '2022-04-10', 'Envolvimento com tráfico para sobreviver', FALSE),

-- Mothers Milk
(9, 9, '2023-02-28', 'Agressão a membros da Vought durante investigação', FALSE),

-- Frenchie
(10, 9, '2022-12-12', 'Produção e uso de drogas ilegais', TRUE),
(10, 6, '2021-05-05', 'Sequestro de criminoso para obter informações', FALSE);

-- Recupera heróis ordenados pela força (descendente) e missões pela dificuldade (descendente)
DO $$
DECLARE
    heroi RECORD;
    missao RECORD;
BEGIN
    -- Itera sobre todas as missões
    FOR missao IN 
        SELECT id_missao, nivel_dificuldade 
        FROM missoes 
        ORDER BY nivel_dificuldade DESC 
    LOOP
        -- Associa os heróis mais fortes para a missão atual
        FOR heroi IN 
            SELECT id_heroi, forca 
            FROM herois 
            WHERE forca >= missao.nivel_dificuldade
            ORDER BY forca DESC 
            LIMIT 3 -- Limita a 3 heróis por missão (ajuste conforme necessário)
        LOOP
            INSERT INTO herois_missoes (id_heroi, id_missao)
            VALUES (heroi.id_heroi, missao.id_missao);
        END LOOP;
    END LOOP;
END $$;
