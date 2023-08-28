


IMageName=benchmark-grpc

Version=v0.1-delay


grpcdelay(){

cat > grpc-delay-10ms.yaml <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: grpc-delay-10ms
  name: grpc-delay-10ms
spec:
  replicas: 1
  selector:
    matchLabels:
      app: grpc-delay-10ms
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: grpc-delay-10ms
    spec:
      containers:
      - image: docker.io/hefabao/$IMageName:$Version
        name: busybox
        resources: {}
        resources:
          limits:
            cpu: 0.25
            memory: 900Mi
          requests:
            cpu: 0.1
            memory: 256Mi
        env:
          - name: "ENV-SERVER-GRPC"
            value: "true"
---
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: null
  labels:
    app: grpc-delay-10ms
  name: grpc-delay-10ms
spec:
  ports:
  - name: p60000
    port: 60000
    protocol: TCP
    targetPort: 80
    nodePort: 30800
  selector:
    app: grpc-delay-10ms
  type: NodePort
---

  
EOF


kubectl apply -f grpc-delay-10ms.yaml


}



httpdelay(){

cat > http-delay.yaml <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: http-delay
  name: http-delay
spec:
  replicas: 1
  selector:
    matchLabels:
      app: http-delay
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: http-delay
    spec:
      containers:
      - image: docker.io/hefabao/$IMageName:$Version
        name: busybox
        resources: {}
        resources:
          limits:
            cpu: 0.25
            memory: 900Mi
          requests:
            cpu: 0.1
            memory: 256Mi
        env:
          - name: "ENV-SERVER-HTTP"
            value: "true"
---
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: null
  labels:
    app: http-delay
  name: http-delay
spec:
  ports:
  - name: p60001
    port: 60001
    protocol: TCP
    targetPort: 60001
    nodePort: 30901
  selector:
    app: http-delay
  type: NodePort
---

  
EOF

kubectl apply -f http-delay.yaml

}


httpdelay

grpcdelay

