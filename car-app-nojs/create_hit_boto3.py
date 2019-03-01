# import boto3
# region_name = 'us-east-1'
# aws_access_key_id = 'AKIAJACQ5XLZN7RPOVRQ'
# aws_secret_access_key = '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'
# endpoint_url = 'https://mturk-requester-sandbox.us-east-1.amazonaws.com'

# # Uncomment this line to use in production
# # endpoint_url = 'https://mturk-requester.us-east-1.amazonaws.com'

import sys
import os
import boto3
import urllib.request
import argparse
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
import smtplib

def send_email(task_number, response, environments, create_hits_in_live):
    server_addr = 'smtp.gmail.com'
    server_port = 587
    user = 'berkeleyvfgteam@gmail.com'
    password = '-t9M7dFY4)/SHi)t'

    mturk_environment = environments["live"] if create_hits_in_live else environments["sandbox"]

    # if sandbox:
    recipient = 'microwhatevermer@gmail.com'
    # else:
        # recipient = 'danielbenniah@googlegroups.com'
    
    hit_type_id = response['HIT']['HITTypeId']
    hit_id = response['HIT']['HITId']

    server = smtplib.SMTP(server_addr, server_port)
    server.ehlo()
    server.starttls()
    server.login(user, password)

    link_msg = mturk_environment['preview'] + "?groupId={}".format(hit_type_id)
    # hit_id_msg = 'HIT ID: {0}'.format(hit_id) 

    subject = 'Task {0} is published'.format(task_number)
    # body = hit_id_msg + '<br>' + link_msg
    body = link_msg
    # print("Body = " + body)
    message = MIMEText(body, 'html')
    message['Subject'] = subject
    message['From'] = user
    message['To'] = recipient

    server.sendmail(user, recipient, message.as_string())
    server.quit()
    print("Hit created")
    print("Hit ID = ", hit_id)

def create_hit(task_number, environments, create_hits_in_live):
    create_hits_in_live = False

    linesInExternalQuestion = 6
    lineToModify = 2

    mturk_environment = environments["live"] if create_hits_in_live else environments["sandbox"]

    # aws_access_key_id = "AKIAJ4U7LFFGGTEHHU7A"
    # aws_secret_access_key = "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"
    aws_access_key_id = "AKIAIV7WZVP3I4YAW7XA"
    aws_secret_access_key = "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"
    client = boto3.client(
      'mturk',
      endpoint_url=mturk_environment['endpoint'],
      region_name='us-east-1',
      aws_access_key_id=aws_access_key_id,
      aws_secret_access_key=aws_secret_access_key,
    )


    # user_balance = client.get_account_balance()

    # print ("Your account balance is {}".format(user_balance['AvailableBalance']))

    # question_sample = urllib.request.urlopen("https://small.localtunnel.me/").read().decode("utf-8") 
    question_sample = open("external_question.xml").read()
    question_sample_split = question_sample.split("\n")
    question_sample_with_id = ""

    for i in range(linesInExternalQuestion):
        if (i != lineToModify):
            question_sample_with_id += question_sample_split[i] + "\n"
        else:
            question_sample_with_id += question_sample_split[i] + task_number

    # print("Question sample = ", question_sample_with_id)

    # print("Type of question = ", type(question_sample))

    worker_requirements = [{
        'QualificationTypeId': '000000000000000000L0',
        'Comparator': 'GreaterThanOrEqualTo',
        'IntegerValues': [80],
        'RequiredToPreview': True,
    }]

    # Create the HIT
    response = client.create_hit(
        MaxAssignments=3,
        LifetimeInSeconds=600,
        AssignmentDurationInSeconds=600,
        Reward=mturk_environment['reward'],
        Title='Classify the following images',
        Keywords='classification, answer, research, 12345678',
        Description='Classify the following images by question.',
        Question=question_sample_with_id
        # QualificationRequirements=worker_requirements,
    )

    return response

def main(task_number):
    environments = {
            "live": {
                "endpoint": "https://mturk-requester.us-east-1.amazonaws.com",
                "preview": "https://www.mturk.com/mturk/preview",
                "manage": "https://requester.mturk.com/mturk/manageHITs",
                "reward": "0.00"
            },
            "sandbox": {
                "endpoint": "https://mturk-requester-sandbox.us-east-1.amazonaws.com",
                "preview": "https://workersandbox.mturk.com/mturk/preview",
                "manage": "https://requestersandbox.mturk.com/mturk/manageHITs",
                "reward": "0.11"
            },
    }
    create_hits_in_live = False
    response = create_hit(task_number, environments, create_hits_in_live)
    send_email(task_number, response, environments, create_hits_in_live)

    # # The response included several fields that will be helpful later
    # hit_type_id = response['HIT']['HITTypeId']
    # hit_id = response['HIT']['HITId']
    # print ("\nCreated HIT: {}".format(hit_id))

    # print ("\nYou can work the HIT here:")
    # print (mturk_environment['preview'] + "?groupId={}".format(hit_type_id))

    # print ("\nAnd see results here:")
    # print (mturk_environment['manage'])

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('--task', type=int, default=-1)
    args = parser.parse_args()

    task = args.task
    if (task == -1 or task == 0):
        print("Please enter a valid task number --task <task_number>");
        exit()

    elif (os.path.isfile('./tasks/task' + str(task) + '.txt') == False):
        print('File ./tasks/task' + str(task) + '.txt does not exist')

    main(str(task))

    
