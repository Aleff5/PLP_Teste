# Documentação da API de Heróis

## Endpoints Disponíveis

### 1. Listar Todos os Heróis
- **Endpoint:** `/`
- **Método:** GET
- **Descrição:** Retorna informações de todos os heróis cadastrados
- **Não requer corpo na requisição**

### 2. Buscar Herói por Nome
- **Endpoint:** `/heroi`
- **Método:** POST
- **Descrição:** Busca um herói específico pelo nome
- **Corpo da Requisição:**
```json
{
    "nome_heroi": "string"
}
```

### 3. Buscar Heróis por Popularidade
- **Endpoint:** `/heroipop`
- **Método:** POST
- **Descrição:** Retorna heróis com base no nível de popularidade
- **Corpo da Requisição:**
```json
{
    "popularidade": int
}
```

### 4. Buscar Heróis por Status
- **Endpoint:** `/heroistatus`
- **Método:** POST
- **Descrição:** Retorna heróis filtrados por status de atividade
- **Corpo da Requisição:**
```json
{
    "status_atividade": "string"
}
```

### 5. Cadastrar Novo Herói
- **Endpoint:** `/heroicadastra`
- **Método:** POST
- **Descrição:** Cadastra um novo herói com seus poderes
- **Corpo da Requisição:**
```json
{
    "heroi": {
        "nome_heroi": "string",
        "nome_real": "string",
        "sexo": "string",         // "Masculino" ou "Feminino"
        "altura": float,          // em metros (ex: 1.80)
        "local_nascimento": "string",
        "data_nascimento": "YYYY-MM-DDT00:00:00Z", // Formato RFC3339,
        "peso": float,            // em kg
        "popularidade": int,      // valor entre 0 e 100
        "forca": int,            // valor entre 0 e 100
        "status_atividade": "string"  // "Ativo", "Banido" ou "Inativo"
    },
    "ids_poderes": [1, 2, 3]     // array com IDs dos poderes existentes
}
```

### 6. Deletar Herói
- **Endpoint:** `/delete`
- **Método:** POST
- **Descrição:** Remove um herói do sistema
- **Corpo da Requisição:**
```json
{
    "id_heroi": int
}
```

### 7. Consultar Crimes por Herói e Severidade
- **Endpoint:** `/heroieseveridadecrime`
- **Método:** POST
- **Descrição:** Busca crimes associados a um herói dentro de um intervalo de severidade
- **Corpo da Requisição:**
```json
{
    "nome_heroi": "string",
    "severidade_minima": int,
    "severidade_maxima": int
}
```

### 8. Consultar Crimes por Herói
- **Endpoint:** `/heroicrime`
- **Método:** POST
- **Descrição:** Retorna todos os crimes associados a um herói específico
- **Corpo da Requisição:**
```json
{
    "nome_heroi": "string"
}
```

### 9. Listar Todos os Poderes
- **Endpoint:** `/poderes`
- **Método:** GET
- **Descrição:** Retorna lista de todos os poderes disponíveis
- **Não requer corpo na requisição**

### 10. Consultar Crimes por Severidade
- **Endpoint:** `/severidadecrime`
- **Método:** POST
- **Descrição:** Busca crimes dentro de um intervalo de severidade
- **Corpo da Requisição:**
```json
{
    "severidade_minima": int,
    "severidade_maxima": int
}
```

### 11. Editar Herói
- **Endpoint:** `/editar`
- **Método:** PUT
- **Descrição:** Atualiza as informações de um herói existente
- **Corpo da Requisição:**
```json
{
    "nome_heroi": "string", // Nome do herói que será editado
    "heroi_atualizado": {
        "nome_heroi": "string",
        "nome_real": "string",
        "sexo": "string", // "Masculino" ou "Feminino"
        "altura": float,
        "local_nascimento": "string",
        "data_nascimento": "YYYY-MM-DDT00:00:00Z",
        "peso": float,
        "popularidade": int, // Valor entre 0 e 100
        "forca": int, // Valor entre 0 e 100
        "status_atividade": "string" // "Ativo", "Banido" ou "Inativo"
    }
}
```

### 12. Consultar Missões por Herói
- **Endpoint:** `/missao`
- **Método:** POST
- **Descrição:** Retorna as missões associadas a um herói específico
- **Corpo da Requisição:**
```json
{
    "nome_heroi": "string"
}
```

### 13. Simular Batalha entre Heróis
- **Endpoint:** `/simularbatalha`
- **Método:** POST
- **Descrição:** Simula uma batalha entre dois heróis, considerando força, popularidade e fatores aleatórios
- **Corpo da Requisição:**
```json
{
    "heroi1": "string",
    "heroi2": "string"
}
```

**Detalhes da Simulação:**
- A força final de cada herói é calculada considerando:
  - Força base do herói
  - Bônus de popularidade (metade da popularidade)
  - Chance de acerto crítico (+20 de força)
  - Fator aleatório (+15 de força, 20% de chance)
  - Moral durante a luta (+5 ou -5 dependendo de quem está ganhando)
- A chance de vitória é proporcional à força final de cada herói
- O vencedor é determinado considerando as forças finais e um fator aleatório