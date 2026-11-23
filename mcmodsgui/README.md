# MCPanelHAOS

Webpanel und Minecraft Java laufen direkt auf demselben Home-Assistant-OS-Gerät. Die Oberfläche und der CurseForge-Import stammen aus Minepanel. Zwei getrennte Profile teilen sich einen Minecraft-Laufzeitplatz und TCP 25565.

Die App benötigt Zugriff auf die Docker-Engine des HAOS-Hosts. Bei deaktiviertem Schutzmodus hat die App über den Socket technisch Zugriff auf den gesamten Host. Die Minecraft-Container selbst erhalten den Socket nicht.
