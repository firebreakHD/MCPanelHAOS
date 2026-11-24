# Änderungen

## 2.1.0

- Optionaler Parallelbetrieb für bis zu zwei Minecraft-Server mit getrennten Host-Ports, etwa 25565 und 25566. Der Ein-Server-Modus bleibt Standard.
- Docker-Speicherlimit bis 10 GiB. Bei 10 GiB startet Java mit 6 GiB Heap und kann bis 8 GiB Heap nutzen.
- Bestehende Serverprofile, Welten und CurseForge-Einstellungen bleiben beim Update erhalten.

## 2.0.1

- Stopp und Serverwechsel warten bei der Ersteinrichtung auf den Start der Minecraft-Laufzeit. Dadurch erreicht das Stoppsignal den Java-Shutdown-Handler.

## 2.0.0

- Ersetzt Minepanel durch DiscoPanel 2.0.15. Add-on-ID und Repository bleiben für das Update gleich.
- Serverprofile mit frei wählbaren Namen und separaten Weltordnern.
- Exklusiver Serverbetrieb mit gemeinsamem Minecraft-Port 25565 und bestätigtem Wechsel über die GUI.
- 4–6 GiB Java-Heap und 5–7 GiB Docker-Speicherlimit ohne zusätzlichen Swap.
- Native CurseForge-Integration mit API-Schlüssel aus den globalen DiscoPanel-Einstellungen.
- HTTPS-Port 18443 für Neuinstallationen.
- Neue DiscoPanel-Datenstruktur. Für Weltübernahme und Rollback siehe [Dokumentation](DOCS.md).
