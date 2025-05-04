# elysionne

Projet Go généré avec le starter CLI

# 🚀 go-starter — Un starter Go moderne & organisé comme Symfony

Un starter Go **prêt à l’emploi**, pensé pour les développeurs PHP/Symfony souhaitant passer à Go.  
Avec une structure claire, une gestion de configuration centralisée, et un générateur CLI pour créer facilement des modèles, handlers, ou migrations.

---

## 📁 Structure du projet

go-starter/
├── cmd/ → Point d'entrée de l'application
│ └── app/
│ └── main.go → Lanceur principal de l'application
├── internal/ → Code applicatif privé
│ ├── models/ → Entités GORM (équivalent Doctrine)
│ ├── handlers/ → Contrôleurs HTTP (équivalent Symfony\Controller)
│ ├── services/ → Logique métier
│ ├── repositories/ → Accès à la base de données
├── config/ → Fichiers de configuration
│ ├── .env → Variables d’environnement
│ └── config.yaml → Configuration structurée (port, logs, etc.)
├── generator/ → Outils CLI pour générer des fichiers
│ └── generator.go → Générateur de modèle/handler/migration
├── migrations/ → Migrations GORM
├── go.mod → Dépendances Go (équivalent composer.json)
└── README.md → Ce fichier


---

## 🧰 Technologies utilisées

| Tech | Rôle |
|------|------|
| [Gin](https://gin-gonic.com/) | Framework web léger |
| [GORM](https://gorm.io/) | ORM pour MySQL/PostgreSQL |
| [GoDotEnv](https://github.com/joho/godotenv) | Chargement des variables `.env` |
| [YAML](https://pkg.go.dev/gopkg.in/yaml.v3) | Lecture de fichiers YAML |

---

## 🛠️ Prérequis

- [Go 1.22+](https://go.dev/dl/)
- Un IDE comme GoLand ou VS Code
- Une base MySQL locale (XAMPP / WAMP / Docker)

---

## 🔽 Installation

1. Clone ce dépôt :
   ```bash
   git clone https://github.com/consolidis/go-starter.git monprojet
   cd monprojet

2. Installe les dépendances 

 > go mod download

3. Configure ta base dans .env

 >  DB_USER=root
    DB_PASSWORD=
    DB_HOST=localhost
    DB_NAME=monsuperprojet
    APP_PORT=8080

4. Lance l’application :

> go run cmd/app/main.go

5. Visite : http://localhost:8080 

🧱 Architecture détaillée

📁 cmd/app/main.go
Point d'entrée de l'application.

 . Charge la configuration
 . Connecte la base de données
 . Démarre le serveur Gin

📁 internal/models/

Contient les structs Go représentant tes tables SQL (comme les entités Doctrine).

Exemple:

type Personnage struct {
    ID   uint   `gorm:"primaryKey"`
    Nom  string
}

📁 internal/handlers/

Équivalent des contrôleurs Symfony. Contient les fonctions gérant les requêtes HTTP.

Exemple:

func GetPersonnage(c *gin.Context) {
    c.JSON(200, gin.H{"message": "Hello from Personnage!"})
}

📁 config/config.yaml

Fichier de configuration structurée :

server:
  port: 8080
log_level: info

📁 generator/generator.go
Outil CLI pour générer automatiquement :

 . Modèles (make:model)
 . Handlers (make:handler)
 . Migrations (make:migration)

 Usage :
 > go run generator/generator.go make:model Personnage
 > go run generator/generator.go make:handler Personnage
 > go run generator/generator.go make:migration Personnage

📁 migrations/

Fichiers pour migrer ta base via GORM.

Exemple :
func MigratePersonnage(db *gorm.DB) {
    db.AutoMigrate(&Personnage{})
}

🧪 Exemple : Créer un CRUD simple

 1. Génère les fichiers 

`go run generator/generator.go make:model Personnage`
`go run generator/generator.go make:handler Personnage`
`go run generator/generator.go make:migration Personnage`

 2. Implémente le handler dans `internal/handlers/personnage_handler.go`.

 3. Lie au routeur dans main.go
 `r.GET("/personnages", handlers.GetPersonnage)`

 3. Lance les migrations manuellement ou automatiquement.

 ------------------------------------------------------------------------------------------------------

 📦 Dépendances : Comment ça marche ?

En Go, les dépendances sont gérées par go.mod.

--------------------------------------------------------------------------------------------------------

🧪 Lancer l’application

Depuis la racine :

 `go run cmd/app/main.go`

 Visiter :
👉 `http://localhost:8080`

-----------------------------------------------------------------------------------------------------------

📦 Construire l’exécutable

 `go build -o myapp cmd/app/main.go`
 `./myapp`
