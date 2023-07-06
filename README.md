# jtt808-gateway

与一般网关不同，该项目是808业务网关，承载808终端（定位设备）和各业务平台之间的通讯。jtt808-gateway 只针对该场景实现了部分网关功能，包含但不限于：

- 【TODO】**安全控制**，实现身份认证、权限验证、黑白名单、IP过滤等功能，保证系统的安全性
- 【TODO】**限流降级**，根据请求的来源、频率、优先级等，实现限流和降级策略，防止系统过载
- 【TODO】**日志监控**，记录请求和响应的日志，监控系统的运行状态和性能指标

## 交互流程
![img.png](http://cdn-0.plantuml.com/plantuml/png/SoWkIImgAStDuKeloYyjK7Y-k-FvwlNFroryFg6DIq71FLp1HbTN8I0diIGjloZNIY6vA3Mn9BLOeIIrA9TB0GYGjJtRlE9fMmzG60H2W2eKT7NjW6POAHIb5fQc5fT0fHAJIpBBWCO0fI0YG0P8ALWFI75nGNvUSIfKBYKLNNrgNWgGXZWAnPBS_B9KYDD0Ib2yzB9poxEvpN0lGQrAB2t9u49r42x7e3a4quQd-vjVRD-CRaD6AZ1yau21HZLhrirwihSNtPgS_EJ4aipyF5os868mCIynfx8uinZ8I05QMv1VL9APbuwa3wuMFVvnkXCkJvAl7804CE410000)

## RTVS 支持
![img.png](http://cdn-0.plantuml.com/plantuml/png/SoWkIImgAStDuKeloYyjK7YwRjwpwTlqd_QrFUtVz69vsyj54vzjRNo-efihA2GiM45Nrqx1FTnAmICa942XABMmDBMuXCiz72mUabgKQwLWK65fQ62e3wIC30ovG69WIP1kAo0Pm9MQby85jUpP_AKlrYzwsZ7znS8LnDY2C2YtE5qXCmN3PYC3KuILxk4WCKEW65nlQafciPM2TsgbbOA_SqIyVv2BQv2ddvj_VFIpdlQdwsPvkiJ6v8KbvfIcMYc4PQQavliNfQGMEIPdlzYxnYlOtm_gF8HHHQc9AQWgmVvp01B0C0K0)

