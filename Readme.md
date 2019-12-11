# <center> <h1> DAYTA SE TEST </h1> </center>

<p align="center">
  <img src="https://dayta-public-assets.s3-ap-southeast-1.amazonaws.com/dayta_logo_full_color.png" height="100"/>
</p>

## GOLANG GRPC AUTH SERVER

### Tasks:

1. Setup the go environment on the host machine.

2. Clone the current repository

3. Implement a go GRPC server that provdides the service described by the service.proto file in /proto folder

4. Fill in the REST API endpoint on the client server. (You can send the token string as an "Authorization" header or in the body. It is up to you.)

5. Dockerize the client & server separately in a alpine(Linux) container

6. Using the client server to send the token, DECODE the following token for next instruction.

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtZXNzYWdlIjoiQ29uZ3JhdHVsYXRpb25zIG9uIGdldHRpbmcgdGhlIGRlY29kZWQgdG9rZW4hIFBsZWFzZSBsZXQgbWUga25vdyBpZiB5b3UndmUgcmVhY2hlZCB0aGlzIHN0ZXAgOikiLCJ0aXRsZSI6IldlbGNvbWUgdG8gRGF5dGEgQUkifQ.VofJqlUwUfWIsVdAb-X9Zz84ApqIQUM8q9CESON2HUA
