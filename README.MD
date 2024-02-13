# How to build and run project:
cd mailgunsender
go build -C cmd/mailgunsender/ -o ../../bin/mailgun-sender
./bin/mailgun-sender -ml test_data/ml.csv -camp DummyCampaign

# .env variables:
MG_API_KEY=
MG_DOMAIN=