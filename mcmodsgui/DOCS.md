# MCPanelHAOS einrichten

Die [Installationsanleitung](../README.md) beschreibt das Repository, den Schutzmodus und die erste Einrichtung. DiscoPanel übernimmt Weboberfläche, Live-Konsole, Dateimanager, CPU-/RAM-Anzeige und Server-Backups. Die HA-Startseite und die Meldungen für den Wechsel sind auf Deutsch. Die übrige DiscoPanel-Oberfläche verwendet die Sprache des Quellprojekts.

## CurseForge

Trage deinen API-Schlüssel in **Global Settings → CurseForge API Key** ein und speichere. Unter **Modpacks** kannst du CurseForge durchsuchen und die Version eines Modpacks auswählen. Bei Ablehnung prüfe, ob der Schlüssel in der [CurseForge-Konsole](https://console.curseforge.com/#/api-keys) aktiv ist. Stelle sicher, dass HAOS das Internet per HTTPS erreicht. Wenn du den Schlüssel bei einem bereits erstellten CurseForge-Server änderst, stoppe diesen und verwende im Panel **Recreate**, damit der Container den aktuellen Schlüssel übernimmt. Die Weltdaten liegen im Serverordner und werden beim Neuerstellen des Containers wieder eingebunden.

## Ports

| Port | Verwendung |
| --- | --- |
| TCP 18443 | DiscoPanel über HTTPS im LAN |
| TCP 25565 | Minecraft Java, für beide Profile dieselbe HAOS-Adresse |
| TCP 443 ausgehend | CurseForge, Download-Quellen und Container-Registries |

Wenn du die externe HTTPS-Portzuordnung in Home Assistant änderst, trage denselben Wert in `panel_port` ein. Für einen anderen DNS-Namen oder eine LAN-IP ändere `panel_hostname`. Über die Startseite des Add-ons kannst du das lokale CA-Zertifikat für deinen Browser herunterladen.

## Datensicherung und Wechsel von Minepanel

Das Add-on speichert Profile, Datenbank, Welten und Panel-Backups im privaten, dauerhaften Ordner `/app/discopanel`. Halte den Zugriff auf vollständige Add-on-Backups beschränkt. Sie enthalten auch private Einstellungen. Der Supervisor stoppt das Add-on vor einem Add-on-Backup. Für einzelne Welten kannst du in DiscoPanel ein Server-Backup anlegen.

Version 2.0.0 verwendet einen neuen Datenordner. Stoppe vor dem Upgrade den Minecraft-Server in Minepanel und sichere das Add-on. Die bisherigen Welten und Profile werden von DiscoPanel nicht eingelesen. Für die Übernahme einer Welt erstellst du in DiscoPanel ein neues Profil mit derselben Minecraft- und Modpack-Version und kopierst die gesicherten Weltdaten bei gestopptem Server über den Dateimanager. Zum Rollback das Add-on stoppen und das Add-on-Backup von vor dem Upgrade wiederherstellen. Sichere zwischenzeitlich erstellte DiscoPanel-Welten vorher separat.

## Betrieb

DiscoPanel nutzt die vom Supervisor freigegebene Docker-API für die Java-Container. Der gemeinsame Port sowie die Startsperre gelten für die von dieser Installation verwalteten Profile. Vom Panel gestartete Automatisierungen können einen bereits laufenden Server nicht verdrängen. Beim bestätigten Wechsel über **Start** wird der vorherige Server regulär beendet. Bei einem Stoppfehler bricht der Wechsel ab.

Der HAOS-Wrapper baut [DiscoPanel 2.0.15](https://github.com/discohaus/discopanel/tree/ba7eb27c9728bb34999f16d77838dfd17a53d191). Lizenz und Herkunft des übernommenen Codes liegen im [Add-on-Ordner](LICENSE).
