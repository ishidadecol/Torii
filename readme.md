<div align="center">

```
████████╗ ██████╗ ██████╗ ██╗██╗
╚══██╔══╝██╔═══██╗██╔══██╗██║██║
   ██║   ██║   ██║██████╔╝██║██║
   ██║   ██║   ██║██╔══██╗██║██║
   ██║   ╚██████╔╝██║  ██║██║██║
   ╚═╝    ╚═════╝ ╚═╝  ╚═╝╚═╝╚═╝
        ⛩  T O R I I  ⛩
```

A lightweight localhost tunneling tool written in **Go**

</div>

---

## Overview

Torii exposes services running on your local machine to the public internet through a TCP tunnel.

> ⚠️ Torii is **heavily inspired by ngrok and slim**.  
> This project exists primarily as a **learning-oriented implementation** to understand how tunneling systems work internally.

```
Internet → Torii Server → Tunnel → Localhost
```

This allows developers to:

- test webhooks locally  
- expose development servers  
- demo applications without deployment  
- learn low-level networking concepts  
