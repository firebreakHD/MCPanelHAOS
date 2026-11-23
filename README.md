# MCPanelHAOS

Ein Home-Assistant-OS-Add-on für zwei Minecraft-Java-Modpacks auf **demselben** HAOS-Rechner. Das Webpanel basiert auf [Minepanel](https://github.com/Ketbome/minepanel). Panel, Minecraft und Welten brauchen keinen weiteren Rechner.

## Das eigene GitHub-Repository veröffentlichen

1. In GitHub Desktop **File → Add Local Repository** öffnen und den Ordner `MCPanelHAOS` wählen.
2. **Publish repository** anklicken, den Namen **MCPanelHAOS** verwenden und **Keep this code private** abwählen. Home Assistant kann das öffentliche Repository ohne GitHub-Login abrufen.
3. Die URL des veröffentlichten Repositories in Home Assistant unter **Einstellungen → Apps/Add-ons → Installieren → ⋮ → Repositories** hinzufügen. **MCPanelHAOS** auswählen und installieren.

## Starten

1. Für **MCPanelHAOS** den Schutzmodus deaktivieren und starten.
2. Die Seitenleiste einblenden und **MCPanelHAOS** öffnen. Beim ersten Aufruf das lokale HTTPS-Stammzertifikat über die Seitenleiste beziehen. Dann die Minecraft-Verwaltung auf demselben HAOS-Gerät öffnen.
3. In der Oberfläche zwei Java-Profile `modpack-a` und `modpack-b` einrichten und CurseForge-Packs auswählen. Spieler verbinden sich mit der HAOS-Adresse auf **25565/TCP**.

[Einrichtung, Ports und Backup](mcmodsgui/DOCS.md) · [Versionshinweise](mcmodsgui/CHANGELOG.md)

## Zugriff und Architektur

Die App nutzt `docker_api: true`. Wie bei Portainer bietet der Socket der App bei ausgeschaltetem Schutzmodus technisch Kontrolle über die Host-Docker-Engine. Die App bekommt keine zusätzlich angeforderten Linux-Capabilities, keine Home-Assistant-API und keine Host-Netzwerkfreigabe. Der Supervisor verwaltet das Panel; das Panel verwaltet seine Minecraft- und Backup-Container auf demselben Host. Der Minecraft-Container bekommt keinen Docker-Socket.

Die App erzeugt 4 GiB maximalen Java-Heap und 5 GiB tatsächliches Container-RAM-Limit. Der Startschutz lässt genau einen Minecraft-Container der App gleichzeitig zu. Der zweite Server startet erst, wenn der vorherige seine Welt gespeichert hat und beendet ist. Die Daten liegen im HAOS-App-Datenmount und bleiben über Updates erhalten. Für einen Produktivtest zuerst getrennte Testwelten verwenden.

Das Repository ist eigenständig. Die GUI und der Servercode basieren auf Minepanel v1.13.44 von Ketbome sowie der lokalen Anpassung für exklusive Profile. Quellcode und Lizenzhinweise liegen vollständig bei. [AGPL-3.0](mcmodsgui/LICENSE) · [Upstream](https://github.com/Ketbome/minepanel)
