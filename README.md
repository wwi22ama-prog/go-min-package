# Go-Projekt: Einfaches Package

## Beschreibung

Das hier gezeigte Beispielprojekt kann als Basis bzw. Kopiervorlage für
ein einfaches Go-Package verwendet werden.
Es gibt hier keine `main()`-Funktion und das Package muss auch nicht `main`
heißen.

Neben der eigentlichen Quelldatei gibt es eine Testdatei, deren Name auf `_test.go`
endet. In dieser Datei kann man Tests schreiben, die die Funktionen aus der eigentlichen
Quelldatei auf Fehler überprüfen.

Dieses Beispielprojekt enthält drei Dateien:

* `sum.go`: Eine Go-Quelldatei, in der eine Funktion `Sum()` definiert wird.
            Diese Funktion wird dort erklärt und implementiert.
* `sum_test.go`: Eine Go-Quelldatei, in der ein Test für die Summenfunktion
                 enthalten ist. Dieser Test ist dazu gedacht, die Korrektheit
                 der Funktion `Sum()` zu prüfen und ihr Verhalten zu dokumentieren.
* `go.mod`: Eine Textdatei, in der Metadaten zum Projekt spezifiziert werden.
            Der Inhalt dieser Datei ist i.W. eine Formalität, mehr dazu unten.

Diese Datei (`README.md`) gehört nicht zum Projekt bzw. sie ist nicht notwendig,
um ein Go-Programm zu schreiben.

## Verwendung

Dieses Minimalprojekt kann als Kopiervorlage benutzt werden.
Es gelten die gleichen Überlegungen und Anweisungen wie beim minimalen
ausführbaren Programm.

In diesem Projekt ist allerdings kein Programm enthalten, das direkt ausgeführt
und vom Benutzer verwendet werden kann.
Stattdessen wird ein *Package* definiert.
Wird dieses Package unter einem geeigneten Namen (z.B. über GitHub) veröffentlicht,
kann es in anderen Go-Projekten mittels eines `import`-Statements eingebunden werden.

Die Namen und Inhalte der beiden Dateien sind natürlich nur Beispiele und können
bei Bedarf geändert werden.

## Erklärung zu go.mod

Diese Datei enthält eine Zeile mit dem Namen des Go-Moduls, das hier
erstellt wird. Falls das Projekt nicht öffentlich ist, spielt dieser Name
keine Rolle, er muss lediglich eindeutig sein.
D.h. falls man einen Workspace mit mehreren Modulen hat, darf dieser Name
nicht mehrfach vorkommen.

Ist das Projekt öffentlich, sollte dieser Modulname der Url entsprechen, unter der das
Projekt auffindbar ist.

Außerdem enthält diese Datei noch die Go-Version, mit der dieses Projekt erstellt wurde.
Falls komplexere Sprach-Features verwendet werden, kann dies bedeuten, dass diese
Go-Version auch benötigt wird, um das Projekt zu verwenden.
