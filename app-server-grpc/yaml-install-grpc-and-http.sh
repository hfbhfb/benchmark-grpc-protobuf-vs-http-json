


yamlinstallgrpc(){

cat > grpc-server.yaml <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: grpc-server
  name: grpc-server
spec:
  replicas: 1
  selector:
    matchLabels:
      app: grpc-server
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: grpc-server
    spec:
      containers:
      - image: docker.io/hefabao/benchmark-grpc:v0.1
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
    app: grpc-server
  name: grpc-server
spec:
  ports:
  - name: p60000
    port: 60000
    protocol: TCP
    targetPort: 60000
    nodePort: 30600
  selector:
    app: grpc-server
  type: NodePort
---

  
EOF


kubectl apply -f grpc-server.yaml


}


yamlinstallgrpc




yamlinstallhttp(){

cat > http-server.yaml <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: http-server
  name: http-server
spec:
  replicas: 1
  selector:
    matchLabels:
      app: http-server
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: http-server
    spec:
      containers:
      - image: docker.io/hefabao/benchmark-grpc:v0.1
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
    app: http-server
  name: http-server
spec:
  ports:
  - name: p60001
    port: 60001
    protocol: TCP
    targetPort: 60001
    nodePort: 30601
  selector:
    app: http-server
  type: NodePort
---

  
EOF

kubectl apply -f http-server.yaml

}


yamlinstallhttp

