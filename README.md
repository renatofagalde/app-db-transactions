# 🧪 Go + GORM – Estudo de Race Condition (Controle de Concorrência)

Este repositório existe **exclusivamente para estudo** de concorrência, race condition e estratégias de correção em aplicações Go utilizando **GORM + PostgreSQL**, mantendo uma **estrutura limpa (inspirada em Clean Architecture)**.

O objetivo aqui **não é performance**, nem ser um exemplo production-ready, e sim **entender claramente o problema e as soluções**, sem ruído desnecessário.

---

## 🎯 Objetivo do projeto

- Demonstrar **race condition real** em cenários concorrentes
- Mostrar que **estrutura limpa não impede bugs de concorrência**
- Comparar **abordagens corretas** sem uso de lock explícito em tabela
- Estudar o comportamento real do **GORM + PostgreSQL**

---

## 🌿 Organização de branches

### 🔴 `feature/01-fail` — **Exemplo com falha (intencional)**

Esta branch **DEMONSTRA O PROBLEMA**.

Características:

- ❌ Não usa transação explícita
- ❌ Não define isolation level
- ❌ Não usa lock pessimista
- ❌ Não usa controle de versão
- ❌ Leitura + atualização separadas
- ✅ Código limpo, organizado e correto *do ponto de vista estrutural*

Resultado:

- Em execução concorrente (ex: 100 goroutines)
- A coluna `quantidade` pode ficar **negativa**
- Resultado final **não determinístico**

👉 Esta branch serve para **ver o bug acontecer**.

---

### 🟢 `main` — **Implementação corrigida (sem lock de tabela)**

Esta branch contém a **correção do problema**, **sem usar lock pessimista** (`FOR UPDATE`).

A solução adotada é:

- ✅ **Optimistic Lock**
- ✅ Coluna `version`
- ✅ Atualização condicional
- ✅ Concorrência controlada pela própria engine do banco

Exemplo conceitual:

```sql
UPDATE estoque
SET quantidade = quantidade - 1,
    version = version + 1
WHERE id = ?
  AND version = ?
```

Se nenhuma linha for afetada, significa que outro processo atualizou antes — evitando a race condition **sem bloquear a tabela ou a linha**.

---

## ⚠️ IMPORTANTE — Sobre `SERIALIZABLE` e isolamento de transação

> **Este projeto NÃO possui (e não precisa possuir) uma branch usando `ISOLATION LEVEL SERIALIZABLE`.**

### Motivo (leia com atenção para não esquecer no futuro):

Quando você usa **GORM** e executa uma operação que:

- realiza um `INSERT` ou `UPDATE`
- e retorna o `ID` do registro
- dentro do mesmo fluxo de execução

O **PostgreSQL já garante serialização implícita daquela operação**, pois:

- o banco precisa garantir que o `ID` retornado é consistente
- isso exige um nível de isolamento suficiente para aquela escrita
- portanto, definir manualmente `SERIALIZABLE` **não muda o comportamento prático** nesse cenário

Em outras palavras:

- Criar uma branch apenas para demonstrar `SERIALIZABLE`
- ❌ não acrescentaria aprendizado real
- ❌ aumentaria complexidade
- ❌ desviaria do foco do estudo

Este repositório **não é sobre isolation level**, e sim sobre:

- ❌ evitar `SELECT ... FOR UPDATE`
- ❌ evitar lock pessimista
- ✅ entender **Optimistic Concurrency Control**
- ✅ entender **race condition real no código**

---

## 🧠 Estrutura do projeto (todas as branches)

```
.
├── cmd
│   └── main.go                  # Simulação concorrente (goroutines)
├── domain
│   └── estoque.go               # Entidade + regra de negócio
├── application
│   └── usecase
│       └── comprar_estoque.go   # Caso de uso
├── infrastructure
│   └── repository
│       └── estoque_gorm.go      # Persistência (GORM)
```

---

## 🧪 Como usar para estudo

1. Checkout da branch com falha:
   ```bash
   git checkout feature/01-fail
   ```

2. Execute o programa
3. Observe o valor final da coluna `quantidade` (pode ficar negativo)

Depois:

4. Volte para a branch principal:
   ```bash
   git checkout main
   ```

5. Execute exatamente o mesmo cenário
6. Observe que:
   - Não há lock explícito
   - Não há isolamento manual
   - A integridade é mantida

---

## 🧠 Nota para o seu “eu do futuro”

> **Estrutura limpa NÃO resolve concorrência.**  
> **Isolation level alto NÃO é bala de prata.**

O que resolve este problema específico é:

- controle otimista
- atualização condicional
- falhar rápido quando há concorrência

Guarde este projeto como **material de referência** para revisar sempre que surgir dúvida sobre race condition, GORM ou concorrência em banco de dados.

---

## 📌 Status

✔ Projeto educacional  
✔ Falha intencional documentada  
✔ Correção real sem lock  
✔ Material de estudo para longo prazo
