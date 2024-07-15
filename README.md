# Projekt Przetwarzania Obrazów

## Przegląd
Ten projekt to aplikacja do przetwarzania obrazów napisana w języku Go. Aplikacja zawiera różne funkcje do wczytywania, przetwarzania oraz zapisywania obrazów. Projekt demonstruje wykorzystanie bibliotek Go do manipulacji obrazami oraz wykonywania operacji na danych obrazowych.

## Funkcje
- **Wczytywanie Obrazów:** Możliwość wczytywania obrazów z różnych formatów.
- **Przetwarzanie Obrazów:** Różne techniki przetwarzania obrazów, takie jak filtrowanie, transformacje oraz efekty graficzne.
- **Zapisywanie Obrazów:** Możliwość zapisywania przetworzonych obrazów do plików.

## Wymagania
- Go (najlepiej najnowsza wersja)
- Biblioteki Go do obsługi obrazów (`image`, `image/color`, `image/draw`, `image/jpeg`, `image/png`)

## Instalacja
1. Sklonuj repozytorium:
    ```sh
    git clone https://github.com/Assassin-PL/GoLang.git
    ```
2. Przejdź do katalogu projektu:
    ```sh
    cd your-repository
    ```
3. Zainstaluj wymagane biblioteki:
    ```sh
    go get -u golang.org/x/image
    ```

## Użycie
1. Uruchom główny plik programu:
    ```sh
    go run MAIN.go
    ```
2. Aplikacja załaduje obraz, wykona przetwarzanie i zapisze wynikowy obraz do pliku.

## Opis Plików
- `MAIN.go`: Główny plik programu, który inicjalizuje i uruchamia aplikację.
- `funkcjePrzetwarzania.go`: Zawiera funkcje do przetwarzania obrazów, takie jak filtrowanie i transformacje.
- `modulWeWy.go`: Zawiera funkcje do wczytywania i zapisywania obrazów.
- `painting.go`: Definiuje funkcje do rysowania i aplikowania efektów graficznych na obrazach.
- `stale.go`: Zawiera stałe używane w całej aplikacji.

## Przykładowe Użycie
Aby uruchomić aplikację, wykonaj następujące kroki:

1. Otwórz terminal i przejdź do katalogu projektu.
2. Uruchom aplikację:
    ```sh
    go run MAIN.go
    ```
3. Aplikacja przetworzy obraz zgodnie z wbudowanymi funkcjami i zapisze wynikowy obraz do pliku.

## Wkład w Projekt
Wkłady są mile widziane! Proszę forkować repozytorium i przesyłać pull requesty.

## Licencja
Ten projekt jest licencjonowany na podstawie licencji MIT. Zobacz plik [LICENSE](LICENSE) po więcej szczegółów.

---

Jeśli napotkasz jakiekolwiek problemy lub masz sugestie dotyczące usprawnień, proszę dać znać. Miłej pracy!
