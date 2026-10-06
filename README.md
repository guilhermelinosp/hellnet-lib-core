# hellnet-lib-core

Centraliza o que e **compartilhado e reutilizado** entre os servicos (fast-platform, fast-sockets, fast-listeners e os proximos).
Se dois servicos precisam do mesmo codigo, ele mora aqui, e nao e copiado de um servico para o outro.

| Pacote | O que e |
|---|---|
| `env` | Carrega `.env` e le variaveis de ambiente |
| `events` | Contratos dos eventos trocados entre os servicos |
| `platform` | Base HTTP (gin), contexto, middleware e ciclo de vida do processo |

```go
import (
    "github.com/guilhermelinosp/hellnet-lib-core/env"
    "github.com/guilhermelinosp/hellnet-lib-core/events"
    "github.com/guilhermelinosp/hellnet-lib-core/platform"
)
```

Regras: so entra aqui o que mais de um servico usa; nada especifico de um servico (regra de negocio fica no servico).
Telemetria, banco, cache e Kafka continuam em suas proprias libs (`hellnet-lib-telemetry`, `-database`, `-cache`, `-kafka`).
Os releases saem sozinhos pelo `pipeline.yml` (semver pelos commits).
