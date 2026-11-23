# MCPanelHAOS

DiscoPanel als Home-Assistant-OS-Add-on für Minecraft Java auf demselben Gerät. Erstelle zwei getrennte Serverprofile mit eigenen Namen, Welten und Modpacks. Es kann jeweils ein Profil laufen. Beide sind unter derselben HAOS-IP über TCP-Port **25565** erreichbar.

## Installation

1. [Repository in Home Assistant hinzufügen](https://my.home-assistant.io/redirect/supervisor_add_addon_repository/?repository_url=https%3A%2F%2Fgithub.com%2FfirebreakHD%2FMCPanelHAOS) oder im Add-on-Store unter **Repositorys** die Adresse `https://github.com/firebreakHD/MCPanelHAOS` eintragen.
2. **MCPanelHAOS** installieren. Bei einem Update von Minepanel zuerst den Minecraft-Server stoppen und ein HA-Backup erstellen.
3. **Schutzmodus** beim Add-on deaktivieren. Die vom Supervisor bereitgestellte Docker-API gibt dem Add-on weitreichenden Zugriff auf den Host. Das Add-on ist für HAOS auf AMD64 gebaut.
4. Falls du statt `homeassistant.local` eine andere Adresse verwendest, trage sie unter **Konfiguration** in `panel_hostname` ein. Der externe HTTPS-Port ist **18443**. Bei einem vorhandenen Setup prüfe **Netzwerk** und `panel_port` in **Konfiguration**. Beide müssen den gleichen externen Port verwenden.
5. Add-on starten. **Weboberfläche öffnen** führt zum DiscoPanel-Login. Die deutsche Startseite in der Home-Assistant-Seitenleiste enthält den Link und das lokale HTTPS-Stammzertifikat.
6. Beim ersten Aufruf einen lokalen Admin erstellen.

## CurseForge und Serverwechsel

Den **CurseForge API Key** in DiscoPanel unter **Global Settings** speichern. Gib dort den Schlüssel aus der [CurseForge-Konsole](https://console.curseforge.com/#/api-keys) ein. Suche danach unter **Modpacks → CurseForge**. Für die Abfragen und Downloads benötigt HAOS ausgehendes HTTPS über TCP 443. Für CurseForge ist keine Portweiterleitung am Router nötig. Manche Mods haben vom Anbieter gesperrte automatische Downloads. DiscoPanel beziehungsweise das Serverprotokoll zeigt dir dann, welche Dateien du manuell hochladen musst. [DiscoPanel-Dokumentation](https://docs.discopanel.app/guides/modpacks/)

Erstelle einen Vanilla-Server unter **Servers → New Server** oder wähle ein Modpack unter **Modpacks**. Neue Server starten standardmäßig erst auf deinen Klick. Wenn ein Server bereits läuft, klickst du beim anderen auf **Start** und bestätigst den Wechsel. Das Panel wartet auf die Beendigung des bisherigen Containers. Wenn beide aus sein sollen, klicke beim laufenden Server auf **Stop**. Lass **Auto Start** für beide aus, wenn nach einem Add-on-Neustart beide aus bleiben sollen.

Der Standard ist **4 GiB Java-Heap** und **5 GiB hartes Docker-Limit** ohne zusätzlichen Swap. Stelle für 5 oder 6 GiB Heap das Docker-Limit am jeweiligen Server auf 6144 oder 7168 MiB. Die zusätzlichen 1024 MiB sind für JVM- und Container-Overhead reserviert.

[Details zu Daten, Backups und Update](mcmodsgui/DOCS.md) · [Änderungen](mcmodsgui/CHANGELOG.md) · [DiscoPanel-Quellprojekt](https://github.com/discohaus/discopanel)
