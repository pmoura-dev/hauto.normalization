# hauto.normalization

`hauto.normalization` provides a normalization layer for the Home Automation system.
Its purpose is to translate between device-specific MQTT messages and the internal conventions used by `hauto`.

The repository defines two core services:

- Ingestor (devices → hauto)
    - Subscribes to device topics
    -  Normalizes incoming messages from various manufacturers
    - Publishes them into the hauto standard topic structure

- Dispatcher (hauto → devices)
    - Subscribes to normalized hauto topics (e.g., commands, state updates)
    - Translates them into the device-specific formats
    - Publishes them back to the appropriate device topics

This architecture ensures that `hauto` always works with a consistent API, while still supporting devices with different MQTT payloads, topic structures, or conventions.