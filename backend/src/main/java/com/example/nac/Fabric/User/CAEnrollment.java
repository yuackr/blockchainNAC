package main.java.com.example.nac.Fabric.User;

import org.hyperledger.fabric.sdk.Enrollment;

import java.io.Serializable;
import java.security.PrivateKey;

public class CAEnrollment implements Enrollment, Serializable {
	private static final long serialVersionUID = 550416591376968096L;
	private PrivateKey key;
	private String cert;

	public CAEnrollment(PrivateKey pkey, String signedPem) {
		this.key = pkey;
		this.cert = signedPem;
	}

	public PrivateKey getKey() {
		return key;
	}

	public String getCert() {
		return cert;
	}
}
