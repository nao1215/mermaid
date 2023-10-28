## Sequence Diagram
```mermaid
sequenceDiagram
    participant Viktoriya
    participant Naohiro
    participant Naoyuki

    Viktoriya->>Naohiro: Please wake up Naoyuki
    Naohiro-->>Viktoriya: OK

    loop until Naoyuki wake up
    Naohiro->>Naoyuki: Wake up!
    Naoyuki-->>Naohiro: zzz
    Naohiro->>Naoyuki: Hey!!!
    break if Naoyuki wake up
    Naoyuki-->>Naohiro: ......
    end
    end

    Naohiro-->>Viktoriya: wake up, wake up
```