package main.java.com.example.nac.Fabric.Util;

import main.java.com.example.nac.Fabric.User.CAEnrollment;
import main.java.com.example.nac.Fabric.User.UserContext;
import org.apache.commons.compress.archivers.tar.TarArchiveEntry;
import org.apache.commons.compress.archivers.tar.TarArchiveOutputStream;
import org.apache.commons.compress.compressors.gzip.GzipCompressorOutputStream;
import org.apache.commons.io.FilenameUtils;
import org.apache.commons.io.IOUtils;
import org.apache.commons.compress.archivers.ArchiveEntry;
import org.hyperledger.fabric.sdk.exception.CryptoException;
import org.hyperledger.fabric.sdk.helper.Utils;

import javax.xml.bind.DatatypeConverter;
import java.io.*;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.security.KeyFactory;
import java.security.NoSuchAlgorithmException;
import java.security.PrivateKey;
import java.security.spec.InvalidKeySpecException;
import java.security.spec.PKCS8EncodedKeySpec;
import java.util.Collection;
import java.util.logging.Level;
import java.util.logging.Logger;

import static java.lang.String.format;


public class Util {

	public static void writeUserContext(UserContext userContext) throws Exception {
		String directoryPath = "users/" + userContext.getAffiliation();
		String filePath = directoryPath + "/" + userContext.getName() + ".ser";
		File directory = new File(directoryPath);
		if (!directory.exists())
			directory.mkdirs();

		FileOutputStream file = new FileOutputStream(filePath);
		ObjectOutputStream out = new ObjectOutputStream(file);

		// Method for serialization of object
		out.writeObject(userContext);

		out.close();
		file.close();
	}

	public static File findFileSk(File directory) {

		File[] matches = directory.listFiles((dir, name) -> name.endsWith("_sk"));

		if (null == matches) {
			throw new RuntimeException(format("Matches returned null does %s directory exist?", directory.getAbsoluteFile().getName()));
		}

		if (matches.length != 1) {
			throw new RuntimeException(format("Expected in %s only 1 sk file but found %d", directory.getAbsoluteFile().getName(), matches.length));
		}

		return matches[0];

	}

	public static InputStream generateTarGzInputStream(File src, String pathPrefix) throws IOException {
		File sourceDirectory = src;

		ByteArrayOutputStream bos = new ByteArrayOutputStream(500000);

		String sourcePath = sourceDirectory.getAbsolutePath();

		TarArchiveOutputStream archiveOutputStream = new TarArchiveOutputStream(new GzipCompressorOutputStream(new BufferedOutputStream(bos)));
		archiveOutputStream.setLongFileMode(TarArchiveOutputStream.LONGFILE_GNU);

		try {
			Collection<File> childrenFiles = org.apache.commons.io.FileUtils.listFiles(sourceDirectory, null, true);

			ArchiveEntry archiveEntry;
			FileInputStream fileInputStream;
			for (File childFile : childrenFiles) {
				String childPath = childFile.getAbsolutePath();
				String relativePath = childPath.substring((sourcePath.length() + 1), childPath.length());

				if (pathPrefix != null) {
					relativePath = Utils.combinePaths(pathPrefix, relativePath);
				}

				relativePath = FilenameUtils.separatorsToUnix(relativePath);

				archiveEntry = new TarArchiveEntry(childFile, relativePath);
				fileInputStream = new FileInputStream(childFile);
				archiveOutputStream.putArchiveEntry(archiveEntry);

				try {
					IOUtils.copy(fileInputStream, archiveOutputStream);
				} finally {
					IOUtils.closeQuietly(fileInputStream);
					archiveOutputStream.closeArchiveEntry();
				}
			}
		} finally {
			IOUtils.closeQuietly(archiveOutputStream);
		}

		return new ByteArrayInputStream(bos.toByteArray());
	}

	public static UserContext readUserContext(String affiliation, String username) throws Exception {
		String filePath = "users/" + affiliation + "/" + username + ".ser";
		File file = new File(filePath);
		if (file.exists()) {
			// Reading the object from a file
			FileInputStream fileStream = new FileInputStream(filePath);
			ObjectInputStream in = new ObjectInputStream(fileStream);

			// Method for deserialization of object
			UserContext uContext = (UserContext) in.readObject();

			in.close();
			fileStream.close();
			return uContext;
		}

		return null;
	}

	public static CAEnrollment getEnrollment(String keyFolderPath,  String keyFileName,  String certFolderPath, String certFileName)
			throws IOException, NoSuchAlgorithmException, InvalidKeySpecException, CryptoException {
		PrivateKey key = null;
		String certificate = null;
		InputStream isKey = null;
		BufferedReader brKey = null;

		try {

			isKey = new FileInputStream(keyFolderPath + File.separator + keyFileName);
			brKey = new BufferedReader(new InputStreamReader(isKey));
			StringBuilder keyBuilder = new StringBuilder();

			for (String line = brKey.readLine(); line != null; line = brKey.readLine()) {
				if (line.indexOf("PRIVATE") == -1) {
					keyBuilder.append(line);
				}
			}

			certificate = new String(Files.readAllBytes(Paths.get(certFolderPath, certFileName)));

			byte[] encoded = DatatypeConverter.parseBase64Binary(keyBuilder.toString());
			PKCS8EncodedKeySpec keySpec = new PKCS8EncodedKeySpec(encoded);
			KeyFactory kf = KeyFactory.getInstance("EC");
			key = kf.generatePrivate(keySpec);
		} finally {
			isKey.close();
			brKey.close();
		}

		CAEnrollment enrollment = new CAEnrollment(key, certificate);
		return enrollment;
	}

	public static void cleanUp() {
		String directoryPath = "users";
		File directory = new File(directoryPath);
		deleteDirectory(directory);
	}

	  public static boolean deleteDirectory(File dir) {
	        if (dir.isDirectory()) {
	            File[] children = dir.listFiles();
	            for (int i = 0; i < children.length; i++) {
	                boolean success = deleteDirectory(children[i]);
	                if (!success) {
	                    return false;
	                }
	            }
	        }

	        // either file or an empty directory
	        Logger.getLogger(Util.class.getName()).log(Level.INFO, "Deleting - " + dir.getName());
	        return dir.delete();
	    }

}
