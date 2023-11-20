Temat projektu: Program do filtrowania obrazów.(np wyostrzanie, uwydatnianie itp). pisany w języku GoLang

uzywane pakiety:
"flag"
	"fmt"
	_ "image/png"
	"os"
	"image"
	"image/jpeg"
	_ "image/png"
	"path/filepath"
    	"github.com/disintegration/imaging"
    "math"
	"math/rand"

Cel projektu: Celem projektu jest stworzenie programu do filtrowania obrazów. Przy tworzeniu aplikacji inspirowalem sie dokumentacja filtrów z programu photoshop i 
stworzylem podobne funkcje, które przetwarzaja dane zdjecie. Dodatkowo skorzystalem z biblioteki "path/filepath" z modulu go aby moc otworzyc i zapisac plik.
Do obslugi zdjecia uzylem biblioteki image, gdzie zapisywalem w zmiennej "image" przykladowy obraz i na tej zmiennej wykonywane byly funkcje ktorych celem bylo
zmienienie wygladu obrazka. Utworzylem tez zmienne w funkcji MAIN ktora, byla glowana funkcja mojego programu, zmienne okreslajace jakie procentowe zmiany maja byc 
przy przetwarzaniu obrazu. do tego stworzylem modul satle.go ktore prechowuje stale w taki sposob aby bylo latow skalowac wplyw zmiennych na program 

Funkcje w programie i jaka role one pelnia:
    Main.go - jest to glwona funkcja mojego programu. W niej zadeklarowalem na poczatku zmienne , ktore okreslaja niebedne parametry do przetwarzania obrazu. Rowniez dodalem flagi ktore maja na celu uporzadkowanie , ktora funkcja przetwarzania ma  byc wywolywana wraz z opisem ktora flaga co robi aby zachowac porzadek w kodzie. Nastepnie utworzylem funkcje changePainting ktora ma na celu nadzoraownie i wykonywanie kazdego polecenia

    funkcjePrzetwarzania.go - w tym module znajduja sie funkcje ktore przetwarzja dana zmienna image, funkcje zostaly zimporoowane z modulu: "github.com/disintegration/imaging" oraz dodalem wlasnorecznie napisana funkcje dodajaca szum Gaussowski do obrazka

    modulWeWy.go - jest to modul odpowiadajcy za wczytywanie obrazka z danego katalogu z obrazami i zapisujacy go. problemem napotkanym w tym module byl problem ze kazde zmiany
    prowadzace przy zmiennej img , moga wywolac nieporzadane amizny na orginalnym obrazie wiec uzylem dunkcji err ktora sprawdza czy dane sa zgodne i czy plik nadaje sie 
    do przypisania do zmiennej img , jesli nie jest wywolywana funkcja zamykania systemu aby program dzialal prawidlowo i chroniaca orginalny plik

    painting.go - jest to funkcja w ktorej nastepuje wczytanie obrazka z podanej lokalizacji i przetwarzanie jego. Na poczatku przelaczam ktore funkcje maja byc aktywne, nastepnie wczytyje z modulWeWy.go plik do zmiennej image i przechodze do zmian jej struktury data ( w ktorej znajduja sie warosci RGb kazdego piskela). Po wykananych zmianach przechodze do zamkniecia pliku 

    stale.go -  w tym module znajduja sie stale , ktorych na celu ma byc ujednolicenie zmiennych aby mozna bylo podawac zmienne w wartosci procentowej, modul jest potrzebny aby 
    uzytkonik nie przekroczyc zbyt duzej wartosci i stabilizowanie programu

Foldery do ktorych zapisujemy obraz znajduja sie w katalogu PrzetwarzanieObrazu
    w katalogu Obrazy znajduja sie obrazy ktore program pobiera i wykonuje nad nimi proces filtrowania
    w katalogu Przetworzone Obrazy\NoweObrazy znajduja sie obrazy ktore zostaly zapisane przez program

Glowne Problemy napotkane podczas pisania kodu i ich rozwiazania.
1. problem z prawidlowym wczytaniem pliku i jego zapisem.
    problem naprawilem tworzac modul modulWeWy ktory sprawdza popraawnosc spisanej sciezki w ktorej znajduje sie obraz, 
2. problem z mozliowscia naruszania orginalnego pliku obrazu
    zastosowalem do rozwiazania tego problemu ze modul modulWeWy pobiera sciezke a zwraca kopie obrazka znajdujacego sie w nim, jest to rozwiazanie kosztowne z 
    punktu uzycia pamieci ale powoduje to ze obraz orginalny nie jest naruszany przez ktorakolwiek z funkcji
3. problem z funkcja szumu Gaussowskiego, a mianowicie jesli plik jest zapisywany w formacie kolorow YCbC i trzeba przekonwertowac na format RGB
    zastosowalem funkcje color.RGBAModel.Convert aby przekonwertowac typ koloru na inny 
4. mialem kolejny problem w funkcji applyDistortion w ktorej rowniez sam edytowalem obraz ale zastosowalem rozwiazania z podpunktu 3.

Mozliwe przyszle zmiany w kodzie:

Kod ma mozliwosc zmiany z poziomu aplikacji z terminala do stworzenia GUI i w czsie rzeczywistym zmiany obrazka, gdzie z lewej strony jest obraz pierwotny a z prawej strony przetworzony. Mozliwoscia zminan jest dodanie checkboxow ktore zmieniaja flagi i suwaka ktory zmienia poziom efektow ( po to w zalozeniu projektu wystepuja falgi i wyskalowanie zmiennych)


